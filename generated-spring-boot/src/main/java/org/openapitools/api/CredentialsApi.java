/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech) (7.8.0).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
package org.openapitools.api;

import org.openapitools.model.BasicErrorResponse;
import org.openapitools.model.CredentialValidationResponse;
import org.openapitools.model.HealthCheckResponse;
import io.swagger.v3.oas.annotations.ExternalDocumentation;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.Parameters;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import io.swagger.v3.oas.annotations.enums.ParameterIn;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.context.request.NativeWebRequest;
import org.springframework.web.multipart.MultipartFile;

import javax.validation.Valid;
import javax.validation.constraints.*;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Generated;

@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", date = "2024-09-25T11:04:41.168990+05:30[Asia/Kolkata]", comments = "Generator version: 7.8.0")
@Validated
@Tag(name = "Authentication", description = "Endpoints to validate the credentials used for calling the Tax Service Adapter.")
public interface CredentialsApi {

    default Optional<NativeWebRequest> getRequest() {
        return Optional.empty();
    }

    /**
     * POST /credentials/validate : Validate credentials
     * This endpoint is used to validate the credentials used to call the Service Provider.
     *
     * @return Authentication succeeded. (status code 200)
     *         or Authentication failed. (status code 401)
     *         or Too many requests. (status code 429)
     *         or Unexpected error while processing request. (status code 500)
     *         or Service is unhealthy. (status code 503)
     */
    @Operation(
        operationId = "validateCredentials",
        summary = "Validate credentials",
        description = "This endpoint is used to validate the credentials used to call the Service Provider.",
        tags = { "Authentication" },
        responses = {
            @ApiResponse(responseCode = "200", description = "Authentication succeeded.", content = {
                @Content(mediaType = "application/json", schema = @Schema(implementation = CredentialValidationResponse.class))
            }),
            @ApiResponse(responseCode = "401", description = "Authentication failed.", content = {
                @Content(mediaType = "application/json", schema = @Schema(implementation = CredentialValidationResponse.class))
            }),
            @ApiResponse(responseCode = "429", description = "Too many requests."),
            @ApiResponse(responseCode = "500", description = "Unexpected error while processing request.", content = {
                @Content(mediaType = "application/json", schema = @Schema(implementation = BasicErrorResponse.class))
            }),
            @ApiResponse(responseCode = "503", description = "Service is unhealthy.", content = {
                @Content(mediaType = "application/json", schema = @Schema(implementation = HealthCheckResponse.class))
            })
        },
        security = {
            @SecurityRequirement(name = "apiKey"),
            @SecurityRequirement(name = "bearerAuth")
        }
    )
    @RequestMapping(
        method = RequestMethod.POST,
        value = "/credentials/validate",
        produces = { "application/json" }
    )
    
    default ResponseEntity<CredentialValidationResponse> validateCredentials(
        
    ) {
        getRequest().ifPresent(request -> {
            for (MediaType mediaType: MediaType.parseMediaTypes(request.getHeader("Accept"))) {
                if (mediaType.isCompatibleWith(MediaType.valueOf("application/json"))) {
                    String exampleString = "{ \"status\" : \"VALID\" }";
                    ApiUtil.setExampleResponse(request, "application/json", exampleString);
                    break;
                }
                if (mediaType.isCompatibleWith(MediaType.valueOf("application/json"))) {
                    String exampleString = "{ \"status\" : \"VALID\" }";
                    ApiUtil.setExampleResponse(request, "application/json", exampleString);
                    break;
                }
                if (mediaType.isCompatibleWith(MediaType.valueOf("application/json"))) {
                    String exampleString = "{ \"helpUrl\" : \"https://apidocs.chargebee.com#Authentication\", \"message\" : \"Invalid Credentials\" }";
                    ApiUtil.setExampleResponse(request, "application/json", exampleString);
                    break;
                }
                if (mediaType.isCompatibleWith(MediaType.valueOf("application/json"))) {
                    String exampleString = "{ \"components\" : [ { \"endpoints\" : [ \"endpoints\", \"endpoints\" ], \"name\" : \"name\", \"description\" : \"description\", \"id\" : \"id\", \"type\" : \"OTHER\" }, { \"endpoints\" : [ \"endpoints\", \"endpoints\" ], \"name\" : \"name\", \"description\" : \"description\", \"id\" : \"id\", \"type\" : \"OTHER\" } ], \"description\" : \"description\", \"time\" : \"2000-01-23T04:56:07.000+00:00\", \"version\" : \"version\", \"status\" : \"UP\" }";
                    ApiUtil.setExampleResponse(request, "application/json", exampleString);
                    break;
                }
            }
        });
        return new ResponseEntity<>(HttpStatus.OK);

    }

}
