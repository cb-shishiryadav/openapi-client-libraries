import unittest

from flask import json

from openapi_server.models.address_validation_request import AddressValidationRequest  # noqa: E501
from openapi_server.models.address_validation_response import AddressValidationResponse  # noqa: E501
from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.check_address_taxability_request import CheckAddressTaxabilityRequest  # noqa: E501
from openapi_server.models.check_address_taxability_response import CheckAddressTaxabilityResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server.test import BaseTestCase


class TestAddressController(BaseTestCase):
    """AddressController integration test stubs"""

    def test_check_address_taxability(self):
        """Test case for check_address_taxability

        Check taxability
        """
        check_address_taxability_request = {"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"}}
        headers = { 
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/address/check-taxability',
            method='POST',
            headers=headers,
            data=json.dumps(check_address_taxability_request),
            content_type='application/json')
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_validate_address(self):
        """Test case for validate_address

        Address validation
        """
        address_validation_request = {"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"}}
        headers = { 
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/address/validate',
            method='POST',
            headers=headers,
            data=json.dumps(address_validation_request),
            content_type='application/json')
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    unittest.main()
