/* tslint:disable */
/* eslint-disable */
/**
 * Feed API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ServerErrorResponse
 */
export interface ServerErrorResponse {
    /**
     * 
     * @type {string}
     * @memberof ServerErrorResponse
     */
    error?: string;
}

/**
 * Check if a given object implements the ServerErrorResponse interface.
 */
export function instanceOfServerErrorResponse(value: object): value is ServerErrorResponse {
    return true;
}

export function ServerErrorResponseFromJSON(json: any): ServerErrorResponse {
    return ServerErrorResponseFromJSONTyped(json, false);
}

export function ServerErrorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ServerErrorResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : json['error'],
    };
}

  export function ServerErrorResponseToJSON(json: any): ServerErrorResponse {
      return ServerErrorResponseToJSONTyped(json, false);
  }

  export function ServerErrorResponseToJSONTyped(value?: ServerErrorResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'error': value['error'],
    };
}

