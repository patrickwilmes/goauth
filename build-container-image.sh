#!/bin/bash

docker_force=false

function check_command() {
  command_to_check=$1
  if ! command -v "$command_to_check" &> /dev/null; then
    echo "$command_to_check is not installed. Please install $command_to_check first."
    exit 1
  fi
}

function check_image_built() {
  container_command=$1
  container_name=$2
  if "$container_command" images | grep -q "$container_name"; then
      echo "Image $container_name successfully built"
      else
        echo "Image could not be created..."
  fi
}

while getopts D opt; do
    case $opt in
        D) docker_force=true ;;
        ?) echo "script usage: $(basename $0) [-D]" >&2; exit 1;;
    esac
done

# remove the parsed options from the positional params
shift $(( OPTIND - 1 ))
container_tag_name=$1

if $docker_force; then
  check_command docker
  docker build -t "$container_tag_name" .
  check_image_built docker "$container_tag_name"
else
  check_command podman
  podman build -t "$container_tag_name" .
  check_image_built podman "$container_tag_name"
fi