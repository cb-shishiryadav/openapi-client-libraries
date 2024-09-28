import unittest

from flask import json

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.invoice import Invoice  # noqa: E501
from openapi_server.models.invoice_request import InvoiceRequest  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server.test import BaseTestCase


class TestInvoiceController(BaseTestCase):
    """InvoiceController integration test stubs"""

    def test_commit_invoice(self):
        """Test case for commit_invoice

        Commit Invoice
        """
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/invoices/{invoice_id}/commit'.format(invoice_id='invoice_id_example'),
            method='POST',
            headers=headers)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_create_invoice(self):
        """Test case for create_invoice

        Create Invoice
        """
        invoice_request = {"seller":{"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"},"taxRegistrationNumber":"taxRegistrationNumber","hasNexus":True},"taxableAmount":5.962133916683182,"discountAmount":1.4658129805029452,"exemptAmount":6.027456183070403,"invoiceCode":"invoiceCode","lineItems":[{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]}],"total":2.3021358869347655,"subtotal":0.8008281904610115,"taxDateTime":"2000-01-23T04:56:07.000+00:00","currency":"currency","taxAmount":5.637376656633329,"documentDateTime":"2000-01-23T04:56:07.000+00:00","customer":{"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"},"name":"name","customerCode":"customerCode","taxRegistrationNumber":"taxRegistrationNumber","hasNexus":True,"locationEvidence":{"paymentCountryCode":"paymentCountryCode","bin":"bin","ip":"ip"},"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]}}
        headers = { 
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/invoices',
            method='POST',
            headers=headers,
            data=json.dumps(invoice_request),
            content_type='application/json')
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_fetch_invoice(self):
        """Test case for fetch_invoice

        Retrieve Invoice
        """
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/invoices/{invoice_id}'.format(invoice_id='invoice_id_example'),
            method='GET',
            headers=headers)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_void_invoice(self):
        """Test case for void_invoice

        Void Invoice
        """
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/invoices/{invoice_id}/void'.format(invoice_id='invoice_id_example'),
            method='POST',
            headers=headers)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    unittest.main()
