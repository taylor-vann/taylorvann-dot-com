# build_and_run_supercache
# brian taylor vann

import os
import json
import subprocess
from string import Template


def create_required_directories():
    if not os.path.exists("cache/conf"):
        os.makedirs("cache/conf")
    if not os.path.exists("cache/data"):
        os.makedirs("cache/data")


def get_config(source):
    config_file = open(source, 'r')
    config = json.load(config_file)
    config_file.close()

    return config


def create_template(source, target, keywords):
    source_file = open(source, 'r')
    source_file_template = Template(source_file.read())
    source_file.close()
    updated_source_file_template = source_file_template.substitute(**keywords)

    target_file = open(target, "w+")
    target_file.write(updated_source_file_template)
    target_file.close()


def create_required_templates(config):
    cache_config = config["cache"]

    create_template("templates/template.dockerfile",
                    "cache/dockerfile", cache_config)

    create_template("templates/redis.template.conf",
                    "cache/conf/redis.conf", cache_config)

    create_template("templates/docker-compose.template.yml",
                    "docker-compose.yml", {"service_name": config["service_name"],
                                           "http_port": config["server"]["http_port"],
                                           "redis_port": cache_config["redis_port"]})


def build_podman_files():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "down"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    create_required_directories()
    config = get_config("config/config.json")
    create_required_templates(config)
    build_podman_files()
