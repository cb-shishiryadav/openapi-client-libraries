import connexion
from typing import Dict
from typing import Tuple
from typing import Union

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.tax_estimation_request import TaxEstimationRequest  # noqa: E501
from openapi_server.models.tax_estimation_response import TaxEstimationResponse  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server import util


def estimate_taxes(tax_estimation_request=None):  # noqa: E501
    """Estimate tax

    This endpoint is used to estimate taxes for a set of line items being sold by the Merchant to a Customer. # noqa: E501

    :param tax_estimation_request: 
    :type tax_estimation_request: dict | bytes

    :rtype: Union[TaxEstimationResponse, Tuple[TaxEstimationResponse, int], Tuple[TaxEstimationResponse, int, Dict[str, str]]
    """
    if connexion.request.is_json:
        tax_estimation_request = TaxEstimationRequest.from_dict(connexion.request.get_json())  # noqa: E501
    return 'do some magic!'
