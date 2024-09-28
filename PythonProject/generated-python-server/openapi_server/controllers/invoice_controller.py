import connexion
from typing import Dict
from typing import Tuple
from typing import Union

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.invoice import Invoice  # noqa: E501
from openapi_server.models.invoice_request import InvoiceRequest  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server import util


def commit_invoice(invoice_id):  # noqa: E501
    """Commit Invoice

    This endpoint is used to commit an invoice for a given invoice id. Once committed, the invoice is considered to be finalized. # noqa: E501

    :param invoice_id: The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.
    :type invoice_id: str

    :rtype: Union[None, Tuple[None, int], Tuple[None, int, Dict[str, str]]
    """
    return 'do some magic!'


def create_invoice(invoice_request=None):  # noqa: E501
    """Create Invoice

    This endpoint is used to send an invoice to the Tax Service Provider. Invoices created in Chargebee are statements of amounts owed by the Customer to the Merchant for a specific purchase. # noqa: E501

    :param invoice_request: 
    :type invoice_request: dict | bytes

    :rtype: Union[Invoice, Tuple[Invoice, int], Tuple[Invoice, int, Dict[str, str]]
    """
    if connexion.request.is_json:
        invoice_request = InvoiceRequest.from_dict(connexion.request.get_json())  # noqa: E501
    return 'do some magic!'


def fetch_invoice(invoice_id):  # noqa: E501
    """Retrieve Invoice

    This endpoint is used to retrieve an invoice for a given invoice id. # noqa: E501

    :param invoice_id: The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.
    :type invoice_id: str

    :rtype: Union[Invoice, Tuple[Invoice, int], Tuple[Invoice, int, Dict[str, str]]
    """
    return 'do some magic!'


def void_invoice(invoice_id):  # noqa: E501
    """Void Invoice

    This endpoint is used to mark a specific invoice as void. Voiding cancels the invoice without deleting it. # noqa: E501

    :param invoice_id: The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.
    :type invoice_id: str

    :rtype: Union[None, Tuple[None, int], Tuple[None, int, Dict[str, str]]
    """
    return 'do some magic!'
