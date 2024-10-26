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
 * @interface ServerCreateResponse
 */
export interface ServerCreateResponse {
    /**
     *
     * @type {string}
     * @memberof ServerCreateResponse
     */
    id?: string;
}

/**
 * Check if a given object implements the ServerCreateResponse interface.
 */
export function instanceOfServerCreateResponse(value: object): value is ServerCreateResponse {
    return true;
}

export function ServerCreateResponseFromJSON(json: any): ServerCreateResponse {
    return ServerCreateResponseFromJSONTyped(json, false);
}

export function ServerCreateResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ServerCreateResponse {
    if (json == null) {
        return json;
    }
    return {

        'id': json['id'] == null ? undefined : json['id'],
    };
}

  export function ServerCreateResponseToJSON(json: any): ServerCreateResponse {
      return ServerCreateResponseToJSONTyped(json, false);
  }

  export function ServerCreateResponseToJSONTyped(value?: ServerCreateResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {

        'id': value['id'],
    };
}
