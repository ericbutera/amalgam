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
import type { ServerUpdateFeed } from './ServerUpdateFeed';
import {
    ServerUpdateFeedFromJSON,
    ServerUpdateFeedFromJSONTyped,
    ServerUpdateFeedToJSON,
    ServerUpdateFeedToJSONTyped,
} from './ServerUpdateFeed';

/**
 * 
 * @export
 * @interface ServerUpdateFeedRequest
 */
export interface ServerUpdateFeedRequest {
    /**
     * 
     * @type {ServerUpdateFeed}
     * @memberof ServerUpdateFeedRequest
     */
    feed?: ServerUpdateFeed;
}

/**
 * Check if a given object implements the ServerUpdateFeedRequest interface.
 */
export function instanceOfServerUpdateFeedRequest(value: object): value is ServerUpdateFeedRequest {
    return true;
}

export function ServerUpdateFeedRequestFromJSON(json: any): ServerUpdateFeedRequest {
    return ServerUpdateFeedRequestFromJSONTyped(json, false);
}

export function ServerUpdateFeedRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ServerUpdateFeedRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'feed': json['feed'] == null ? undefined : ServerUpdateFeedFromJSON(json['feed']),
    };
}

  export function ServerUpdateFeedRequestToJSON(json: any): ServerUpdateFeedRequest {
      return ServerUpdateFeedRequestToJSONTyped(json, false);
  }

  export function ServerUpdateFeedRequestToJSONTyped(value?: ServerUpdateFeedRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'feed': ServerUpdateFeedToJSON(value['feed']),
    };
}

