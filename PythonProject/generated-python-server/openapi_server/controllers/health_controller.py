import connexion
from typing import Dict
from typing import Tuple
from typing import Union

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server import util


def fetch_health():  # noqa: E501
    """Fetch the health status of the Service Provider.

    This endpoint is used to fetch the health status of the Service Provider. # noqa: E501


    :rtype: Union[HealthCheckResponse, Tuple[HealthCheckResponse, int], Tuple[HealthCheckResponse, int, Dict[str, str]]
    """
    return 'do some magic!'
