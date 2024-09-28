import connexion
from typing import Dict
from typing import Tuple
from typing import Union

from openapi_server.models.address_validation_request import AddressValidationRequest  # noqa: E501
from openapi_server.models.address_validation_response import AddressValidationResponse  # noqa: E501
from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.check_address_taxability_request import CheckAddressTaxabilityRequest  # noqa: E501
from openapi_server.models.check_address_taxability_response import CheckAddressTaxabilityResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server import util


def check_address_taxability(check_address_taxability_request=None):  # noqa: E501
    """Check taxability

    Checks whether the tax address is valid in terms of tax calculation. This endpoint checks whether the address information of the customer is sufficient for the tax provider to return a tax rate. It does not consider the nexus status of the merchant and is mandatory to integrate for the tax provider. # noqa: E501

    :param check_address_taxability_request: 
    :type check_address_taxability_request: dict | bytes

    :rtype: Union[CheckAddressTaxabilityResponse, Tuple[CheckAddressTaxabilityResponse, int], Tuple[CheckAddressTaxabilityResponse, int, Dict[str, str]]
    """
    if connexion.request.is_json:
        check_address_taxability_request = CheckAddressTaxabilityRequest.from_dict(connexion.request.get_json())  # noqa: E501
    return 'do some magic!'


def validate_address(address_validation_request=None):  # noqa: E501
    """Address validation

    Checks whether a given address is a valid delivery address for shipping purposes. The tax provider can decide whether to mention the full or valid address depending on their requirement. # noqa: E501

    :param address_validation_request: 
    :type address_validation_request: dict | bytes

    :rtype: Union[AddressValidationResponse, Tuple[AddressValidationResponse, int], Tuple[AddressValidationResponse, int, Dict[str, str]]
    """
    if connexion.request.is_json:
        address_validation_request = AddressValidationRequest.from_dict(connexion.request.get_json())  # noqa: E501
    return 'do some magic!'
