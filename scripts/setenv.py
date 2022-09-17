#!/usr/bin/env python3

import yaml
import json
import argparse
from sys import stderr
from os import path


def read_env_json(filename: str, key: str) -> dict:
    """
    Reads the env file in json format and gets the env_file
    names

        Parameter:
        ----------
            filename (str): The docker-compose file name

        Returns:
        --------
            data (dict): The key parsed from the json file.

    """
    with open(filename, "r") as stream:
        try:
            data: dict = json.load(stream)
            return data[key]
        except json.JSONError as exc:
            print(exc, file=stderr)


def get_env_docker_compose(filename: str, keys: str) -> str:
    """
    Reads the docker compose file and gets the env_file
    names

        Parameter:
        ----------
            filename (str): The docker-compose file name
            keys     (str): The services that need's to be
                            parsed.

        Returns:
        --------
            data (str): The env_file's name with the services
                        as the key.

    """
    with open(filename, "r") as stream:
        try:
            services: dict = yaml.safe_load(stream)['services']
            data: dict = {}

            for service in keys:
                # If service not found in the services
                # then raise an exception.
                try:
                    assert service in services
                    data[service] = services[service]['env_file']
                except AssertionError as e:
                    print(
                        f"{e}\nKey Not Found {service} in the {filename}",
                        file=stderr
                    )
                    exit(-1)

            return data
        except yaml.YAMLError as exc:
            print(exc, file=stderr)
            exit(-1)


def create_env_file(env_files: dict, env_data: dict):
    """
    Creates envirment file for every service defined in the docker-compose.

        Parameter:
        ----------
            env_files (dict): The env file names parsed from the
                              docker-compose.
            env_data  (dict): The env variables loaded from the env json file.

        Returns:
        --------
            res      (bool): If the the file creation operation fails then
                             return False else return True.
    """

    # looping through all the env_files while creating them and
    # populating them with env_data.
    for service, env_file in env_files.items():

        try:
            # Opening/creating file in "w" mode. This will
            # remove any content from the existing file.
            env_file = env_file[0]
            with open(env_file, "w") as e_file:
                for key, value in env_data[service].items():
                    print(f"{key}={value}", file=e_file)
            print(f"[CREATED] {env_file} for {service}.")
        except Exception as e:
            print(e, file=stderr)
            exit(-1)


def create_compose_env_file(env_filename: str, env_data: dict) -> None:
    """
    Creates a global envirment file for the docker-compose file.
        Parameter:
        ----------
            env_filename (str) : The name of the env file.
            env_data    (dict) : The env variables loaded from
                                 the env json file.

        Returns:
        --------
            res      (bool): If the the file creation operation fails then
                             return False else return True.
    """

    try:
        # Opening/creating file in "w" mode. This will
        # remove any content from the existing file.
        with open(env_filename, "w") as e_file:
            for key, value in env_data.items():
                print(f"{key}={value}", file=e_file)
            print(f"[CREATED] {env_filename} for docker-compose.")
    except Exception as e:
        print(e, file=stderr)
        exit(-1)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
                        description='Reads the envirment variables and ' \
                                    'creates separate env files for all the ' \
                                    'services defined in the docker-compose'
            )
    parser.add_argument('FILE', type=str,
                        help='The target docker-compose file.')

    parser.add_argument("DIR", type=str,
                        help="The directory where the processed " \
                             "env files will be placed."
                        )

    parser.add_argument("ENV_JSON", type=str,
                        help="The targeted key env file in json format."
                        )

    parser.add_argument("-t", "--target-key", type=str,
                        help="The targeted key which should be parsed."
                        )

    args = parser.parse_args()

    # loading the env file as a dict.
    data_env: dict = read_env_json(args.ENV_JSON, args.target_key)

    # loading the env file name of every service from the docker-compose.
    docker_env: dict = get_env_docker_compose(args.FILE, data_env.keys())

    # finally create the separated env files.
    create_env_file(docker_env, data_env)

    # Now create one global env file for the docker-compose variables.
    docker_compose_env_file_path: str = \
        path.join(args.DIR, ".env." + args.target_key)

    all_env = {}
    for value in data_env.values():
        all_env = {**all_env, **value}

    create_compose_env_file(docker_compose_env_file_path, all_env)
