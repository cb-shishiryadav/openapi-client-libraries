import sys
from setuptools import setup, find_packages

NAME = "openapi_server"
VERSION = "1.0.0"

# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = [
    "connexion>=2.0.2",
    "swagger-ui-bundle>=0.0.2",
    "python_dateutil>=2.6.0"
]

setup(
    name=NAME,
    version=VERSION,
    description="Taxes Service Adapter SPI",
    author_email="",
    url="",
    keywords=["OpenAPI", "Taxes Service Adapter SPI"],
    install_requires=REQUIRES,
    packages=find_packages(),
    package_data={'': ['openapi/openapi.yaml']},
    include_package_data=True,
    entry_points={
        'console_scripts': ['openapi_server=openapi_server.__main__:main']},
    long_description="""\
    ## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here&#39;s a list of terms we&#39;ve used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider&#39;s API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object &#x60;api_configuration&#x60; required for configuring your onboarding on Chargebee&#39;s marketplace. In the &#x60;api_configuration&#x60; object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from &#x60;credential_configuration.id&#x60; and creates the HTTP header-based input query parameter for authorization. The &#x60;credential_configuration&#x60; is an array of objects with an &#x60;id&#x60; attribute, and the value of &#x60;id&#x60; is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     &#x60;&#x60;&#x60;json       \&quot;api_configuration\&quot;: {           \&quot;api_base_url\&quot;: \&quot;https://xyz.abc.com/chargebee\&quot;,           \&quot;credential_configuration\&quot;: [              {               \&quot;id\&quot;:\&quot;authorization_key\&quot;,               \&quot;name\&quot;: \&quot;Authorization Key\&quot;,                \&quot;type\&quot;: \&quot;text\&quot;,               \&quot;is_sensitive\&quot;: true              },             {               \&quot;id\&quot;: \&quot;client_secret\&quot;,               \&quot;name\&quot;: \&quot;Client Secret\&quot;,               \&quot;type\&quot;: \&quot;text\&quot;,               \&quot;is_sensitive\&quot;: true             }           ]         }    &#x60;&#x60;&#x60;     &#x60;\&quot;Authorization\&quot;:\&quot;{\&quot;&lt;authorization_key&gt;\&quot;:\&quot;test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\&quot;}\&quot;&#x60; 
    """
)

