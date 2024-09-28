import unittest

from flask import json

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.credential_validation_response import CredentialValidationResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.test import BaseTestCase


class TestAuthenticationController(BaseTestCase):
    """AuthenticationController integration test stubs"""

    def test_validate_credentials(self):
        """Test case for validate_credentials

        Validate credentials
        """
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/credentials/validate',
            method='POST',
            headers=headers)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    unittest.main()
