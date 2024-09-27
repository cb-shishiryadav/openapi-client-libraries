# coding: utf-8

"""
    Taxes Service Adapter SPI

    ## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here's a list of terms we've used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider's API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object `api_configuration` required for configuring your onboarding on Chargebee's marketplace. In the `api_configuration` object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from `credential_configuration.id` and creates the HTTP header-based input query parameter for authorization. The `credential_configuration` is an array of objects with an `id` attribute, and the value of `id` is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     ```json       \"api_configuration\": {           \"api_base_url\": \"https://xyz.abc.com/chargebee\",           \"credential_configuration\": [              {               \"id\":\"authorization_key\",               \"name\": \"Authorization Key\",                \"type\": \"text\",               \"is_sensitive\": true              },             {               \"id\": \"client_secret\",               \"name\": \"Client Secret\",               \"type\": \"text\",               \"is_sensitive\": true             }           ]         }    ```     `\"Authorization\":\"{\"<authorization_key>\":\"test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\"}\"` 

    The version of the OpenAPI document: 0.3.7
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from openapi_client.models.invoice import Invoice

class TestInvoice(unittest.TestCase):
    """Invoice unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> Invoice:
        """Test Invoice
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `Invoice`
        """
        model = Invoice()
        if include_optional:
            return Invoice(
                invoice_id = '',
                invoice_code = '',
                document_date_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                tax_date_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                status = 'PENDING',
                currency = '012',
                seller = openapi_client.models.seller.Seller(
                    tax_registration_number = '', 
                    address = openapi_client.models.address.Address(
                        line1 = '', 
                        line2 = '', 
                        line3 = '', 
                        city = '', 
                        state = '', 
                        postal_code = '', 
                        country = '', ), 
                    has_nexus = True, ),
                customer = openapi_client.models.customer.Customer(
                    name = '', 
                    customer_code = '', 
                    address = openapi_client.models.address.Address(
                        line1 = '', 
                        line2 = '', 
                        line3 = '', 
                        city = '', 
                        state = '', 
                        postal_code = '', 
                        country = '', ), 
                    tax_registration_number = '', 
                    tax_identifiers = [
                        openapi_client.models.field_item.FieldItem(
                            id = '', 
                            value = '', )
                        ], 
                    has_nexus = True, 
                    location_evidence = openapi_client.models.customer_location_evidence.CustomerLocationEvidence(
                        ip = '', 
                        bin = '', 
                        payment_country_code = '', ), ),
                subtotal = 1.337,
                exempt_amount = 1.337,
                discount_amount = 1.337,
                taxable_amount = 1.337,
                tax_amount = 1.337,
                total = 1.337,
                line_items = [
                    openapi_client.models.invoice_line_item.InvoiceLineItem(
                        number = 1, 
                        item_code = '', 
                        description = '', 
                        quantity = 0, 
                        unit_price = 0, 
                        amount = 1.337, 
                        subtotal = 1.337, 
                        is_tax_inclusive = True, 
                        is_taxable = True, 
                        tax_identifiers = [
                            openapi_client.models.field_item.FieldItem(
                                id = '', 
                                value = '', )
                            ], 
                        tax_exempt_type = 'PRODUCT_EXEMPT', 
                        tax_exempt_reason = 'The customer is exempt from taxes.', 
                        exempt_amount = 1.337, 
                        discount_amount = 1.337, 
                        taxable_amount = 1.337, 
                        tax_amount = 1.337, 
                        total = 1.337, 
                        is_partial_tax = True, 
                        taxes = [
                            openapi_client.models.tax_line_item.TaxLineItem(
                                number = 1, 
                                jurisdiction = openapi_client.models.tax_jurisdiction.TaxJurisdiction(
                                    code = '', 
                                    type = 'COUNTRY', 
                                    name = '', ), 
                                name = 'GST', 
                                rate = 1.337, 
                                taxable_amount = 1.337, 
                                tax_amount = 1.337, )
                            ], )
                    ]
            )
        else:
            return Invoice(
                invoice_id = '',
                invoice_code = '',
                document_date_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                status = 'PENDING',
                currency = '012',
                seller = openapi_client.models.seller.Seller(
                    tax_registration_number = '', 
                    address = openapi_client.models.address.Address(
                        line1 = '', 
                        line2 = '', 
                        line3 = '', 
                        city = '', 
                        state = '', 
                        postal_code = '', 
                        country = '', ), 
                    has_nexus = True, ),
                customer = openapi_client.models.customer.Customer(
                    name = '', 
                    customer_code = '', 
                    address = openapi_client.models.address.Address(
                        line1 = '', 
                        line2 = '', 
                        line3 = '', 
                        city = '', 
                        state = '', 
                        postal_code = '', 
                        country = '', ), 
                    tax_registration_number = '', 
                    tax_identifiers = [
                        openapi_client.models.field_item.FieldItem(
                            id = '', 
                            value = '', )
                        ], 
                    has_nexus = True, 
                    location_evidence = openapi_client.models.customer_location_evidence.CustomerLocationEvidence(
                        ip = '', 
                        bin = '', 
                        payment_country_code = '', ), ),
                subtotal = 1.337,
                exempt_amount = 1.337,
                discount_amount = 1.337,
                taxable_amount = 1.337,
                tax_amount = 1.337,
                total = 1.337,
                line_items = [
                    openapi_client.models.invoice_line_item.InvoiceLineItem(
                        number = 1, 
                        item_code = '', 
                        description = '', 
                        quantity = 0, 
                        unit_price = 0, 
                        amount = 1.337, 
                        subtotal = 1.337, 
                        is_tax_inclusive = True, 
                        is_taxable = True, 
                        tax_identifiers = [
                            openapi_client.models.field_item.FieldItem(
                                id = '', 
                                value = '', )
                            ], 
                        tax_exempt_type = 'PRODUCT_EXEMPT', 
                        tax_exempt_reason = 'The customer is exempt from taxes.', 
                        exempt_amount = 1.337, 
                        discount_amount = 1.337, 
                        taxable_amount = 1.337, 
                        tax_amount = 1.337, 
                        total = 1.337, 
                        is_partial_tax = True, 
                        taxes = [
                            openapi_client.models.tax_line_item.TaxLineItem(
                                number = 1, 
                                jurisdiction = openapi_client.models.tax_jurisdiction.TaxJurisdiction(
                                    code = '', 
                                    type = 'COUNTRY', 
                                    name = '', ), 
                                name = 'GST', 
                                rate = 1.337, 
                                taxable_amount = 1.337, 
                                tax_amount = 1.337, )
                            ], )
                    ],
        )
        """

    def testInvoice(self):
        """Test Invoice"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
