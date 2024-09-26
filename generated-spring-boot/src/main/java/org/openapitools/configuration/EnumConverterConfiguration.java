package org.openapitools.configuration;

import org.openapitools.model.AddressValidationStatus;
import org.openapitools.model.CredentialStatus;
import org.openapitools.model.CreditNoteType;
import org.openapitools.model.DocumentStatus;
import org.openapitools.model.ErrorCode;
import org.openapitools.model.HealthStatus;
import org.openapitools.model.TaxExemptType;
import org.openapitools.model.TaxJurisdictionType;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.convert.converter.Converter;

@Configuration
public class EnumConverterConfiguration {

    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.addressValidationStatusConverter")
    Converter<String, AddressValidationStatus> addressValidationStatusConverter() {
        return new Converter<String, AddressValidationStatus>() {
            @Override
            public AddressValidationStatus convert(String source) {
                return AddressValidationStatus.fromValue(source);
            }
        };
    }
    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.credentialStatusConverter")
    Converter<String, CredentialStatus> credentialStatusConverter() {
        return new Converter<String, CredentialStatus>() {
            @Override
            public CredentialStatus convert(String source) {
                return CredentialStatus.fromValue(source);
            }
        };
    }
    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.creditNoteTypeConverter")
    Converter<String, CreditNoteType> creditNoteTypeConverter() {
        return new Converter<String, CreditNoteType>() {
            @Override
            public CreditNoteType convert(String source) {
                return CreditNoteType.fromValue(source);
            }
        };
    }
    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.documentStatusConverter")
    Converter<String, DocumentStatus> documentStatusConverter() {
        return new Converter<String, DocumentStatus>() {
            @Override
            public DocumentStatus convert(String source) {
                return DocumentStatus.fromValue(source);
            }
        };
    }
    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.errorCodeConverter")
    Converter<String, ErrorCode> errorCodeConverter() {
        return new Converter<String, ErrorCode>() {
            @Override
            public ErrorCode convert(String source) {
                return ErrorCode.fromValue(source);
            }
        };
    }
    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.healthStatusConverter")
    Converter<String, HealthStatus> healthStatusConverter() {
        return new Converter<String, HealthStatus>() {
            @Override
            public HealthStatus convert(String source) {
                return HealthStatus.fromValue(source);
            }
        };
    }
    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.taxExemptTypeConverter")
    Converter<String, TaxExemptType> taxExemptTypeConverter() {
        return new Converter<String, TaxExemptType>() {
            @Override
            public TaxExemptType convert(String source) {
                return TaxExemptType.fromValue(source);
            }
        };
    }
    @Bean(name = "org.openapitools.configuration.EnumConverterConfiguration.taxJurisdictionTypeConverter")
    Converter<String, TaxJurisdictionType> taxJurisdictionTypeConverter() {
        return new Converter<String, TaxJurisdictionType>() {
            @Override
            public TaxJurisdictionType convert(String source) {
                return TaxJurisdictionType.fromValue(source);
            }
        };
    }

}
