package org.openapitools.model;

import java.net.URI;
import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonTypeName;
import java.net.URI;
import org.openapitools.jackson.nullable.JsonNullable;
import java.time.OffsetDateTime;
import javax.validation.Valid;
import javax.validation.constraints.*;
import io.swagger.v3.oas.annotations.media.Schema;


import java.util.*;
import javax.annotation.Generated;

/**
 * The basic error response containing the error message and the help documentation link.
 */

@Schema(name = "BasicErrorResponse_1", description = "The basic error response containing the error message and the help documentation link.")
@JsonTypeName("BasicErrorResponse_1")
@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", date = "2024-09-25T11:04:41.168990+05:30[Asia/Kolkata]", comments = "Generator version: 7.8.0")
public class BasicErrorResponse1 {

  private String message;

  private URI helpUrl;

  public BasicErrorResponse1() {
    super();
  }

  /**
   * Constructor with only required parameters
   */
  public BasicErrorResponse1(String message) {
    this.message = message;
  }

  public BasicErrorResponse1 message(String message) {
    this.message = message;
    return this;
  }

  /**
   * The description of the error with details about it's cause.
   * @return message
   */
  @NotNull @Size(max = 250) 
  @Schema(name = "message", example = "Invalid Credentials", description = "The description of the error with details about it's cause.", requiredMode = Schema.RequiredMode.REQUIRED)
  @JsonProperty("message")
  public String getMessage() {
    return message;
  }

  public void setMessage(String message) {
    this.message = message;
  }

  public BasicErrorResponse1 helpUrl(URI helpUrl) {
    this.helpUrl = helpUrl;
    return this;
  }

  /**
   * The link to the documentation for more information about the error and the corrective action.
   * @return helpUrl
   */
  @Valid 
  @Schema(name = "helpUrl", example = "https://apidocs.chargebee.com#Authentication", description = "The link to the documentation for more information about the error and the corrective action.", requiredMode = Schema.RequiredMode.NOT_REQUIRED)
  @JsonProperty("helpUrl")
  public URI getHelpUrl() {
    return helpUrl;
  }

  public void setHelpUrl(URI helpUrl) {
    this.helpUrl = helpUrl;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BasicErrorResponse1 basicErrorResponse1 = (BasicErrorResponse1) o;
    return Objects.equals(this.message, basicErrorResponse1.message) &&
        Objects.equals(this.helpUrl, basicErrorResponse1.helpUrl);
  }

  @Override
  public int hashCode() {
    return Objects.hash(message, helpUrl);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BasicErrorResponse1 {\n");
    sb.append("    message: ").append(toIndentedString(message)).append("\n");
    sb.append("    helpUrl: ").append(toIndentedString(helpUrl)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}

