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
import type { GormDeletedAt } from './GormDeletedAt';
import {
    GormDeletedAtFromJSON,
    GormDeletedAtFromJSONTyped,
    GormDeletedAtToJSON,
    GormDeletedAtToJSONTyped,
} from './GormDeletedAt';
import type { ModelsFeed } from './ModelsFeed';
import {
    ModelsFeedFromJSON,
    ModelsFeedFromJSONTyped,
    ModelsFeedToJSON,
    ModelsFeedToJSONTyped,
} from './ModelsFeed';

/**
 * 
 * @export
 * @interface ModelsArticle
 */
export interface ModelsArticle {
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    authorEmail?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    authorName?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    content?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    createdAt?: string;
    /**
     * 
     * @type {GormDeletedAt}
     * @memberof ModelsArticle
     */
    deletedAt?: GormDeletedAt;
    /**
     * 
     * @type {ModelsFeed}
     * @memberof ModelsArticle
     */
    feed?: ModelsFeed;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    feedId: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    guid?: string;
    /**
     * ID        uint           `gorm:"primarykey" json:"id" binding:"required" example:"1"`
     * @type {string}
     * @memberof ModelsArticle
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    imageUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    preview?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    updatedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelsArticle
     */
    url: string;
}

/**
 * Check if a given object implements the ModelsArticle interface.
 */
export function instanceOfModelsArticle(value: object): value is ModelsArticle {
    if (!('feedId' in value) || value['feedId'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function ModelsArticleFromJSON(json: any): ModelsArticle {
    return ModelsArticleFromJSONTyped(json, false);
}

export function ModelsArticleFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelsArticle {
    if (json == null) {
        return json;
    }
    return {
        
        'authorEmail': json['authorEmail'] == null ? undefined : json['authorEmail'],
        'authorName': json['authorName'] == null ? undefined : json['authorName'],
        'content': json['content'] == null ? undefined : json['content'],
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : GormDeletedAtFromJSON(json['deletedAt']),
        'feed': json['feed'] == null ? undefined : ModelsFeedFromJSON(json['feed']),
        'feedId': json['feedId'],
        'guid': json['guid'] == null ? undefined : json['guid'],
        'id': json['id'],
        'imageUrl': json['imageUrl'] == null ? undefined : json['imageUrl'],
        'preview': json['preview'] == null ? undefined : json['preview'],
        'title': json['title'] == null ? undefined : json['title'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'url': json['url'],
    };
}

  export function ModelsArticleToJSON(json: any): ModelsArticle {
      return ModelsArticleToJSONTyped(json, false);
  }

  export function ModelsArticleToJSONTyped(value?: ModelsArticle | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'authorEmail': value['authorEmail'],
        'authorName': value['authorName'],
        'content': value['content'],
        'createdAt': value['createdAt'],
        'deletedAt': GormDeletedAtToJSON(value['deletedAt']),
        'feed': ModelsFeedToJSON(value['feed']),
        'feedId': value['feedId'],
        'guid': value['guid'],
        'id': value['id'],
        'imageUrl': value['imageUrl'],
        'preview': value['preview'],
        'title': value['title'],
        'updatedAt': value['updatedAt'],
        'url': value['url'],
    };
}

