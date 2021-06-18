# build_and_run_gateway
# brian taylor vann

import json
import subprocess
from string import Template


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
    server_config = config["server"]

    create_template("templates/template.dockerfile",
                    "webapi/dockerfile", server_config)

    create_template("templates/docker-compose.template.yml",
                    "docker-compose.yml", {"service_name": config["service_name"],
                                           "http_port": server_config["http_port"],
                                           "https_port": server_config["https_port"]})


def build_and_run_podman():
    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "down"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "build"])

    subprocess.run(["podman-compose", "--file",
                   "./docker-compose.yml", "up", "--detach"])


if __name__ == "__main__":
    config = get_config("config/config.json")
    create_required_templates(config)
    build_and_run_podman()
