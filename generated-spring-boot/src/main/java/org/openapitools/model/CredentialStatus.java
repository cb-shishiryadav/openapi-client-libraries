package org.openapitools.model;

import java.net.URI;
import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonValue;
import org.openapitools.jackson.nullable.JsonNullable;
import java.time.OffsetDateTime;
import javax.validation.Valid;
import javax.validation.constraints.*;
import io.swagger.v3.oas.annotations.media.Schema;


import java.util.*;
import javax.annotation.Generated;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/**
 * The status of the credentials used for authentication.
 */

@Generated(value = "org.openapitools.codegen.languages.SpringCodegen", date = "2024-09-25T11:04:41.168990+05:30[Asia/Kolkata]", comments = "Generator version: 7.8.0")
public enum CredentialStatus {
  
  VALID("VALID"),
  
  INVALID("INVALID");

  private String value;

  CredentialStatus(String value) {
    this.value = value;
  }

  @JsonValue
  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  @JsonCreator
  public static CredentialStatus fromValue(String value) {
    for (CredentialStatus b : CredentialStatus.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}

