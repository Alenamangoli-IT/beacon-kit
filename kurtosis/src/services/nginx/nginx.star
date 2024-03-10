NAME_ARG = "name"
IMAGE_ARG = "image"
PORT_ID = "port_id"
PORT_NUMBER = "port_number"
CONFIG_FILES_ARTIFACT = "nginx_config"
ROOT_DIRPATH = "root_dirpath"
ROOT_FILE_ARTIFACT = "root_file_artifact_name"
DEFAULT_SERVICE_NAME = "nginx"
DEFAULT_IMAGE = "nginx:latest"
DEFAULT_CONFIG_LOCAL_FILEPATH = "./default.conf.template"
DEFAULT_CONFIG_FILEPATH = "/etc/nginx/templates"
DEFAULT_PORT_ID = "http"
DEFAULT_PORT_NUMBER = 80
HTTP_PORT_APP_PROTOCOL = "http"


def get_config(plan, services, args = {}):
    name = args.get(NAME_ARG, DEFAULT_SERVICE_NAME)
    image = args.get(IMAGE_ARG, DEFAULT_IMAGE)
    port_id = args.get(PORT_ID, DEFAULT_PORT_ID)
    port_number = args.get(PORT_NUMBER, DEFAULT_PORT_NUMBER)
    root_dirpath = args.get(ROOT_DIRPATH, "")
    root_file_artifact = args.get(ROOT_FILE_ARTIFACT, "")

    config_file_artifact = plan.upload_files(
        DEFAULT_CONFIG_LOCAL_FILEPATH,
        CONFIG_FILES_ARTIFACT
    )

    files = {
        DEFAULT_CONFIG_FILEPATH: config_file_artifact
    }
        
    if root_dirpath != "" and root_file_artifact != "":
        files[root_dirpath] = root_file_artifact

    formatted_services = []
    for service in services:
        service = "    server {};".format(service)
        formatted_services.append(service)
    load_balanced_services = '''
'''.join(formatted_services)

    plan.print(load_balanced_services)

    nginx_service = plan.add_service(
        name = name,
        config = ServiceConfig(
            image = image,
            ports = {
                port_id: PortSpec(number = port_number, application_protocol = HTTP_PORT_APP_PROTOCOL)
            },
            env_vars = {
                "LOAD_BALANCED_SERVICES": load_balanced_services,
            },
            files = files,
        )
    )

    return nginx_service