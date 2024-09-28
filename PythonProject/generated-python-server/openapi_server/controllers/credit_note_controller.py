import connexion
from typing import Dict
from typing import Tuple
from typing import Union

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.credit_note import CreditNote  # noqa: E501
from openapi_server.models.credit_note_request import CreditNoteRequest  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server import util


def commit_credit_note(credit_note_id, invoice_id=None):  # noqa: E501
    """Commit credit note

    This endpoint is used to mark a credit note as committed. Once committed, the credit note is considered as finalized. # noqa: E501

    :param credit_note_id: The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider.
    :type credit_note_id: str
    :param invoice_id: The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.
    :type invoice_id: str

    :rtype: Union[None, Tuple[None, int], Tuple[None, int, Dict[str, str]]
    """
    return 'do some magic!'


def create_credit_note(credit_note_request=None):  # noqa: E501
    """Create credit note

    This endpoint is used to send a credit note to the Tax Service Adapter. A credit note is used to reduce the amount due on an invoice. If the credit note is issued after payments have been made for the invoice, refunds can be issued to the Customer. # noqa: E501

    :param credit_note_request: 
    :type credit_note_request: dict | bytes

    :rtype: Union[CreditNote, Tuple[CreditNote, int], Tuple[CreditNote, int, Dict[str, str]]
    """
    if connexion.request.is_json:
        credit_note_request = CreditNoteRequest.from_dict(connexion.request.get_json())  # noqa: E501
    return 'do some magic!'


def fetch_credit_note(credit_note_id, invoice_id=None):  # noqa: E501
    """Retrieve credit note

    This endpoint is used to retrieve a specific credit note using the unique credit note id. # noqa: E501

    :param credit_note_id: The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider.
    :type credit_note_id: str
    :param invoice_id: The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.
    :type invoice_id: str

    :rtype: Union[CreditNote, Tuple[CreditNote, int], Tuple[CreditNote, int, Dict[str, str]]
    """
    return 'do some magic!'


def void_credit_note(credit_note_id, invoice_id=None):  # noqa: E501
    """Void credit note

    This endpoint is used to void the credit note for a specific credit note id. Voiding reverses the credit note, thereby restoring the amount due on the invoice. # noqa: E501

    :param credit_note_id: The unique identifier of the credit note at the Tax Service Adapter or Tax Service Provider.
    :type credit_note_id: str
    :param invoice_id: The unique identifier of the invoice at the Tax Service Adapter or Tax Service Provider.
    :type invoice_id: str

    :rtype: Union[None, Tuple[None, int], Tuple[None, int, Dict[str, str]]
    """
    return 'do some magic!'
