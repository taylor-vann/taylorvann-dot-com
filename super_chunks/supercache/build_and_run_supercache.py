import json
import subprocess
from string import Template


def get_config(source):
    config_file = open(source, 'r')
    config = json.load(config_file)
    config_file.close()

    return config


def create_template(source, target, keywords):
    dockerfile = open(source, 'r')
    dockerfile_template = Template(dockerfile.read())
    dockerfile.close()
    updated_dockerfile_template = dockerfile_template.substitute(**keywords)

    dest_dockerfile = open(target, "w")
    dest_dockerfile.write(updated_dockerfile_template)
    dest_dockerfile.close()


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
    config = get_config("config/config.json")
    create_required_templates(config)
    build_podman_files()
