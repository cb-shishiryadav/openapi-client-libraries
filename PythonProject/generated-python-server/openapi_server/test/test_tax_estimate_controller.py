import unittest

from flask import json

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.tax_estimation_request import TaxEstimationRequest  # noqa: E501
from openapi_server.models.tax_estimation_response import TaxEstimationResponse  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server.test import BaseTestCase


class TestTaxEstimateController(BaseTestCase):
    """TaxEstimateController integration test stubs"""

    def test_estimate_taxes(self):
        """Test case for estimate_taxes

        Estimate tax
        """
        tax_estimation_request = {"seller":{"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"},"taxRegistrationNumber":"taxRegistrationNumber","hasNexus":True},"lineItems":[{"unitPrice":0.14658129805029452,"number":1,"amount":5.962133916683182,"quantity":0.6027456183070403,"itemCode":"itemCode","description":"description","discountAmount":5.637376656633329,"isTaxInclusive":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.14658129805029452,"number":1,"amount":5.962133916683182,"quantity":0.6027456183070403,"itemCode":"itemCode","description":"description","discountAmount":5.637376656633329,"isTaxInclusive":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.14658129805029452,"number":1,"amount":5.962133916683182,"quantity":0.6027456183070403,"itemCode":"itemCode","description":"description","discountAmount":5.637376656633329,"isTaxInclusive":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.14658129805029452,"number":1,"amount":5.962133916683182,"quantity":0.6027456183070403,"itemCode":"itemCode","description":"description","discountAmount":5.637376656633329,"isTaxInclusive":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.14658129805029452,"number":1,"amount":5.962133916683182,"quantity":0.6027456183070403,"itemCode":"itemCode","description":"description","discountAmount":5.637376656633329,"isTaxInclusive":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]}],"estimateDateTime":"2000-01-23T04:56:07.000+00:00","currency":"currency","customer":{"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"},"name":"name","customerCode":"customerCode","taxRegistrationNumber":"taxRegistrationNumber","hasNexus":True,"locationEvidence":{"paymentCountryCode":"paymentCountryCode","bin":"bin","ip":"ip"},"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]}}
        headers = { 
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/tax-estimate',
            method='POST',
            headers=headers,
            data=json.dumps(tax_estimation_request),
            content_type='application/json')
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    unittest.main()
