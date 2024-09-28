import unittest

from flask import json

from openapi_server.models.basic_error_response import BasicErrorResponse  # noqa: E501
from openapi_server.models.credit_note import CreditNote  # noqa: E501
from openapi_server.models.credit_note_request import CreditNoteRequest  # noqa: E501
from openapi_server.models.health_check_response import HealthCheckResponse  # noqa: E501
from openapi_server.models.validation_error_response import ValidationErrorResponse  # noqa: E501
from openapi_server.test import BaseTestCase


class TestCreditNoteController(BaseTestCase):
    """CreditNoteController integration test stubs"""

    def test_commit_credit_note(self):
        """Test case for commit_credit_note

        Commit credit note
        """
        query_string = [('invoiceId', 'invoice_id_example')]
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/credit-notes/{credit_note_id}/commit'.format(credit_note_id='credit_note_id_example'),
            method='POST',
            headers=headers,
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_create_credit_note(self):
        """Test case for create_credit_note

        Create credit note
        """
        credit_note_request = {"seller":{"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"},"taxRegistrationNumber":"taxRegistrationNumber","hasNexus":True},"creditNoteType":"FULL","taxableAmount":5.962133916683182,"roundingAmount":2.3021358869347655,"creditNoteCode":"creditNoteCode","discountAmount":1.4658129805029452,"exemptAmount":6.027456183070403,"invoiceCode":"invoiceCode","lineItems":[{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]},{"unitPrice":0.20271230230023218,"taxableAmount":1.4894159098541704,"amount":4.145608029883936,"quantity":0.3616076749251911,"itemCode":"itemCode","description":"description","isTaxable":True,"discountAmount":1.0246457001441578,"taxes":[{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562},{"number":1,"taxableAmount":5.025004791520295,"rate":49.65218492984954,"jurisdiction":{"code":"code","name":"name","type":"COUNTRY"},"name":"GST","taxAmount":9.965781217890562}],"exemptAmount":1.2315135367772556,"taxExemptType":"PRODUCT_EXEMPT","number":1,"total":7.457744773683766,"subtotal":7.386281948385884,"isTaxInclusive":True,"taxAmount":6.84685269835264,"taxExemptReason":"The customer is exempt from taxes.","isPartialTax":True,"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]}],"total":0.8008281904610115,"taxDateTime":"2000-01-23T04:56:07.000+00:00","invoiceId":"invoiceId","currency":"currency","taxAmount":5.637376656633329,"documentDateTime":"2000-01-23T04:56:07.000+00:00","customer":{"address":{"country":"country","city":"city","postalCode":"postalCode","state":"state","line3":"line3","line2":"line2","line1":"line1"},"name":"name","customerCode":"customerCode","taxRegistrationNumber":"taxRegistrationNumber","hasNexus":True,"locationEvidence":{"paymentCountryCode":"paymentCountryCode","bin":"bin","ip":"ip"},"taxIdentifiers":[{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"},{"id":"id","value":"value"}]}}
        headers = { 
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/credit-notes',
            method='POST',
            headers=headers,
            data=json.dumps(credit_note_request),
            content_type='application/json')
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_fetch_credit_note(self):
        """Test case for fetch_credit_note

        Retrieve credit note
        """
        query_string = [('invoiceId', 'invoice_id_example')]
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/credit-notes/{credit_note_id}'.format(credit_note_id='credit_note_id_example'),
            method='GET',
            headers=headers,
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))

    def test_void_credit_note(self):
        """Test case for void_credit_note

        Void credit note
        """
        query_string = [('invoiceId', 'invoice_id_example')]
        headers = { 
            'Accept': 'application/json',
            'apiKey': 'special-key',
            'Authorization': 'Bearer special-key',
        }
        response = self.client.open(
            '/api/v1/credit-notes/{credit_note_id}/void'.format(credit_note_id='credit_note_id_example'),
            method='POST',
            headers=headers,
            query_string=query_string)
        self.assert200(response,
                       'Response body is : ' + response.data.decode('utf-8'))


if __name__ == '__main__':
    unittest.main()
