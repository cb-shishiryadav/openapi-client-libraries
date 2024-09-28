import connexion
from typing import Dict
from typing import Tuple
from typing import Union

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.credential_validation_response import CredentialValidationResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server import util


def validate_credentials():  # noqa: E501
    """Validate credentials

    This endpoint is used to validate the credentials used to call the Service Provider. # noqa: E501


    :rtype: Union[CredentialValidationResponse, Tuple[CredentialValidationResponse, int], Tuple[CredentialValidationResponse, int, Dict[str, str]]
    """
    return 'do some magic!'
