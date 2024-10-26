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
 * @interface ServerListFeed
 */
export interface ServerListFeed {
    /**
     *
     * @type {string}
     * @memberof ServerListFeed
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof ServerListFeed
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof ServerListFeed
     */
    url?: string;
}

/**
 * Check if a given object implements the ServerListFeed interface.
 */
export function instanceOfServerListFeed(value: object): value is ServerListFeed {
    return true;
}

export function ServerListFeedFromJSON(json: any): ServerListFeed {
    return ServerListFeedFromJSONTyped(json, false);
}

export function ServerListFeedFromJSONTyped(json: any, ignoreDiscriminator: boolean): ServerListFeed {
    if (json == null) {
        return json;
    }
    return {

        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

  export function ServerListFeedToJSON(json: any): ServerListFeed {
      return ServerListFeedToJSONTyped(json, false);
  }

  export function ServerListFeedToJSONTyped(value?: ServerListFeed | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {

        'id': value['id'],
        'name': value['name'],
        'url': value['url'],
    };
}