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
 * @interface ServiceArticle
 */
export interface ServiceArticle {
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    authorEmail?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    authorName?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    content?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    feedId?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    guid?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    imageUrl?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    preview?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    title?: string;
    /**
     *
     * @type {string}
     * @memberof ServiceArticle
     */
    url?: string;
}

/**
 * Check if a given object implements the ServiceArticle interface.
 */
export function instanceOfServiceArticle(value: object): value is ServiceArticle {
    return true;
}

export function ServiceArticleFromJSON(json: any): ServiceArticle {
    return ServiceArticleFromJSONTyped(json, false);
}

export function ServiceArticleFromJSONTyped(json: any, ignoreDiscriminator: boolean): ServiceArticle {
    if (json == null) {
        return json;
    }
    return {

        'authorEmail': json['authorEmail'] == null ? undefined : json['authorEmail'],
        'authorName': json['authorName'] == null ? undefined : json['authorName'],
        'content': json['content'] == null ? undefined : json['content'],
        'feedId': json['feedId'] == null ? undefined : json['feedId'],
        'guid': json['guid'] == null ? undefined : json['guid'],
        'id': json['id'] == null ? undefined : json['id'],
        'imageUrl': json['imageUrl'] == null ? undefined : json['imageUrl'],
        'preview': json['preview'] == null ? undefined : json['preview'],
        'title': json['title'] == null ? undefined : json['title'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

  export function ServiceArticleToJSON(json: any): ServiceArticle {
      return ServiceArticleToJSONTyped(json, false);
  }

  export function ServiceArticleToJSONTyped(value?: ServiceArticle | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {

        'authorEmail': value['authorEmail'],
        'authorName': value['authorName'],
        'content': value['content'],
        'feedId': value['feedId'],
        'guid': value['guid'],
        'id': value['id'],
        'imageUrl': value['imageUrl'],
        'preview': value['preview'],
        'title': value['title'],
        'url': value['url'],
    };
}
