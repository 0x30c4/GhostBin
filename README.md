<h1 align="center">GhostBin</h1>
<p align="center">
  <img src="./assets/logo.png">
</p>
<p align="center">
  <img src="https://img.shields.io/github/languages/top/0x30c4/GhostBin?style=flat-square" alt="Test">
</p>
<p align="justify">
  GhostBin is a lightweight, high-performance pastebin built with Go and Redis. Designed with simplicity and speed in mind, GhostBin offers a fast and efficient platform for sharing text snippets effortlessly.
</p>
<p align="center">
  <a href="mailto:support@gbin.me"> support@gbin.me </a>
</p>

<br>
<img src="./assets/ghostbindemo.gif">
<br>

## Table of Contents

- [How To Use GhostBin](#how-to-use-ghostbin)
  - [Command Line Cient](#cli-client)
  - [Basic Usage](#basic-usage)
  - [Pipe Output of a Command](#pipe-output-of-a-command)
  - [Burn and Expire](#burn-and-expire)
  - [Secure Deletion](#secure-deletion)
  - [Example](#example)
- [Deployment](#deployment)
  - [Built With](#built-with)
  - [Prerequisites](#prerequisites)
  - [Docker Compose](#docker-compose)
  - [Contributing](#contributing)
  - [License](#license)
- [Test](#test)
  - [Test Coverage](#current-test-coverage)
- [TODO](#todo)
- [Contribution](#contribution)
- [Donate](#donate)

## How To Use GhostBin

### Command Line Client

You can use this script to upload your pastes easily. <a href="https://raw.githubusercontent.com/0x30c4/GhostBin/main/gbin.sh"> gbin.sh </a>

```bash
Usage: gbin.sh [-f filename] [-e expire_seconds] [-r max_reads] [-d deepurl_length] [-s secret]
  -f: The filename to upload. Use '-' to pipe output of a command.
  -e: Expire time in seconds (default: no expiration).
  -r: Maximum number of reads (default: unlimited).
  -d: Length of the URL (default: random URL length).
  -s: Secret for deletion (default: none).
```

```bash

#!/bin/bash

# Usage: ./ghostbin.sh [-f filename] [-e expire_seconds] [-r max_reads] [-d deepurl_length] [-s secret]

# Default values
EXPIRE="18446744073709551615"  # No expiration by default
READS="0"                      # No read limit by default
DEEPURL="0"                    # Default URL length
SECRET=""                      # No secret by default
FILENAME="-"                   # Default to pipe input

# Function to display usage
usage() {
  echo "Usage: $0 [-f filename] [-e expire_seconds] [-r max_reads] [-d deepurl_length] [-s secret]"
  echo "  -f: The filename to upload. Use '-' to pipe output of a command."
  echo "  -e: Expire time in seconds (default: no expiration)."
  echo "  -r: Maximum number of reads (default: unlimited)."
  echo "  -d: Length of the URL (default: random URL length)."
  echo "  -s: Secret for deletion (default: none)."
  exit 1
}

# Parse command-line arguments
while getopts "f:e:r:d:s:" opt; do
  case ${opt} in
    f ) FILENAME=$OPTARG ;;
    e ) EXPIRE=$OPTARG ;;
    r ) READS=$OPTARG ;;
    d ) DEEPURL=$OPTARG ;;
    s ) SECRET=$OPTARG ;;
    * ) usage ;;
  esac
done

# Check if filename is provided or input is piped
if [ -z "$FILENAME" ] && [ -t 0 ]; then
  usage
fi

# Build the curl command
CURL_CMD="curl -F \"f=@${FILENAME}\""
[ "$EXPIRE" != "18446744073709551615" ] && CURL_CMD+=" -F \"expire=${EXPIRE}\""
[ "$READS" != "0" ] && CURL_CMD+=" -F \"read=${READS}\""
[ "$DEEPURL" != "0" ] && CURL_CMD+=" -F \"deepurl=${DEEPURL}\""
[ -n "$SECRET" ] && CURL_CMD+=" -F \"secret=${SECRET}\""

# Append the GhostBin URL
CURL_CMD+=" gbin.me"

# Execute the curl command
if [ "$FILENAME" == "-" ]; then
  # Handle piped input
  cat | eval $CURL_CMD
else
  # Handle file input
  eval $CURL_CMD
fi
```

### Basic Usage

```bash
# Upload a file
$ curl -F "f=@filename.ext" gbin.me
```

### Pipe Output of a Command

```bash
# Pipe output of a command
$ cat file | curl -F "f=@-" gbin.me
$ find /var/log/nginx -name "*.log" | curl -F "f=@-" gbin.me
```

### Burn and Expire

```bash
# Paste will expire after 69 seconds
$ curl -F "f=@filename.ext" -F "expire=69" gbin.me

# Paste will expire after 3 reads
$ curl -F "f=@filename.ext" -F "read=3" gbin.me

# Set a custom URL length
$ curl -F "f=@filename.ext" -F "deepurl=3" gbin.me
```

### Secure Deletion

```bash
# Set a secret for deletion
$ curl -F "f=@filename.ext" -F "secret=password" gbin.me

# Delete paste using secret
$ curl -XDELETE -F "secret=password" gbin.me/pasteid
```

### Example

```bash
# Create a paste with specific settings
$ curl -F "f=@filename.ext" -F "deepurl=12" -F "expire=69" -F "read=1" gbin.me
```

For more details and advanced usage, please refer to the [documentation](https://gbin.me).

## Deployment
Want to run a server like this? clone it! Remember centralization is bad.

### Built With.

* [Docker](https://www.docker.com) - Platform and Software Deployment
* [Go](https://go.dev) - Backend Frame-work.
* [Redis](https://redis.io/) - DataStore DataStore

### Prerequisites.

Make sure you have [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git), [make](https://tldp.org/HOWTO/Software-Building-HOWTO-3.html) and [Docker](https://www.docker.com/products/docker-desktop) installed.

### Docker Compose

GhostBin can be easily deployed using Docker Compose. Follow these steps to deploy GhostBin:

1. **Clone Repository**: Clone the GhostBin repository to your server.

2. **Configuration**: Duplicate the `env-example` file and rename it as `.env.dev` for local development or `.env.prod` for the production environment. Customize the contents of these files according to your requirements.

3. **Build**: Build the Docker images for GhostBin using the provided Makefile command:

    ```bash
    make build
    ```

4. **Development Environment**:

    ```bash
    make up-dev
    ```

5. **Production Environment**:

    ```bash
    make up-prod
    ```

6. **Access Logs**:

    To access logs, you can use:

    ```bash
    make logs
    ```

    To tail logs in real-time:

    ```bash
    make logs-tail
    ```

7. **Additional Commands**:

    - `make down-dev` / `make down-prod`: Shutdown the development/production environment.
    - `make restart-dev` / `make restart-prod`: Restart the development/production environment.
    - `make exec-dev` / `make exec-prod`: Access the shell of the development/production container.

## Test
I am presently working on writing the unit tests. 🫠

### Current Test Coverage
<img src="./assets/testcover.svg">

## TODO
  - Write test for handlers
  - Write file delete daemon

## Contributing

Contributions to GhostBin are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on our [GitHub repository](https://github.com/0x30c4/GhostBin).

## License

GhostBin is licensed under the [BSD 3-Clause License](https://github.com/0x30c4/GhostBin/blob/main/LICENSE).


## Contribution
Pull requests are welcome.

For major changes, please open an issue first to discuss what you would like to change.

## Donate
You can support this project via Liberapay.
The monthly hosting cost is right now 12 Dollar.
<br>
<a target="_blank" href="https://liberapay.com/sanaf/donate"><img src="https://img.shields.io/liberapay/gives/1"></a>

Monero wallet address: 83BDAy6tN99PVud2sUnjyoMzsUDdXJCoMjjwJ59cVwPF91RccxLWCVsfD9imMqxUaMhMG1brzuVBeAM4KREUSf9U9efbKx1
