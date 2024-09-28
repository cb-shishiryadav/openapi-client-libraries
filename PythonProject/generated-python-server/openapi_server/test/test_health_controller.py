import unittest

from flask import json

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.test import BaseTestCase


class TestHealthController(BaseTestCase):
    """HealthController integration test stubs"""

    def test_fetch_health(self):
        """Test case for fetch_health

        Fetch the health status of the Service Provider.
        """
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/health',
            method='GET',
            headers=headers)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    unittest.main()
