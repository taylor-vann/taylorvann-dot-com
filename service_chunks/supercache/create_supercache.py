import json
from string import Template

config = open('config/config.json', 'r')
config_data = json.load(config)
config.close()

# dockerfile
dockerfile = open('templates/template.dockerfile', 'r')
dockerfile_template = Template(dockerfile.read())
dockerfile.close()
updated_dockerfile_template = dockerfile_template.substitute(config_data["cache"])

dest_dockerfile = open("cache/dockerfile", "w")
dest_dockerfile.write(updated_dockerfile_template)
dest_dockerfile.close()

# redis
redis = open('templates/redis.template.conf', 'r')
redis_template = Template(redis.read())
redis.close()
updated_redis_template = redis_template.substitute(config_data["cache"])

dest_redis = open("cache/conf/redis.conf", "w")
dest_redis.write(updated_redis_template)
dest_redis.close()

# docker-compose
dockercompose = open('templates/docker-compose.template.yml', 'r')
dockercompose_template = Template(dockercompose.read())
dockercompose.close()

http_port = config_data["server"]["http_port"]
redis_port = config_data["cache"]["redis_port"]
updated_dockercompose_template = dockercompose_template.substitute(http_port=http_port, redis_port=redis_port)

dest_dockercompose = open("docker-compose.yml", "w")
dest_dockercompose.write(updated_dockercompose_template)
dest_dockercompose.close()
