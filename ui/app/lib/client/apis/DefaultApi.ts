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


import * as runtime from '../runtime';
import type {
  ServerArticleResponse,
  ServerErrorResponse,
  ServerFeedArticlesResponse,
  ServerFeedCreateResponse,
  ServerFeedResponse,
  ServerFeedUpdateResponse,
  ServerFeedsResponse,
} from '../models/index';
import {
    ServerArticleResponseFromJSON,
    ServerArticleResponseToJSON,
    ServerErrorResponseFromJSON,
    ServerErrorResponseToJSON,
    ServerFeedArticlesResponseFromJSON,
    ServerFeedArticlesResponseToJSON,
    ServerFeedCreateResponseFromJSON,
    ServerFeedCreateResponseToJSON,
    ServerFeedResponseFromJSON,
    ServerFeedResponseToJSON,
    ServerFeedUpdateResponseFromJSON,
    ServerFeedUpdateResponseToJSON,
    ServerFeedsResponseFromJSON,
    ServerFeedsResponseToJSON,
} from '../models/index';

export interface ArticlesIdGetRequest {
    id: number;
}

export interface FeedsIdArticlesGetRequest {
    id: number;
}

export interface FeedsIdGetRequest {
    id: number;
}

export interface FeedsIdPostRequest {
    id: number;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * view article
     * view article
     */
    async articlesIdGetRaw(requestParameters: ArticlesIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ServerArticleResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling articlesIdGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/articles/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ServerArticleResponseFromJSON(jsonValue));
    }

    /**
     * view article
     * view article
     */
    async articlesIdGet(requestParameters: ArticlesIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ServerArticleResponse> {
        const response = await this.articlesIdGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * list feeds
     * list feeds
     */
    async feedsGetRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ServerFeedsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/feeds`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ServerFeedsResponseFromJSON(jsonValue));
    }

    /**
     * list feeds
     * list feeds
     */
    async feedsGet(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ServerFeedsResponse> {
        const response = await this.feedsGetRaw(initOverrides);
        return await response.value();
    }

    /**
     * list articles for a feed
     * list articles for a feed
     */
    async feedsIdArticlesGetRaw(requestParameters: FeedsIdArticlesGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ServerFeedArticlesResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling feedsIdArticlesGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/feeds/{id}/articles`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ServerFeedArticlesResponseFromJSON(jsonValue));
    }

    /**
     * list articles for a feed
     * list articles for a feed
     */
    async feedsIdArticlesGet(requestParameters: FeedsIdArticlesGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ServerFeedArticlesResponse> {
        const response = await this.feedsIdArticlesGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * view feed
     * view feed
     */
    async feedsIdGetRaw(requestParameters: FeedsIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ServerFeedResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling feedsIdGet().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/feeds/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ServerFeedResponseFromJSON(jsonValue));
    }

    /**
     * view feed
     * view feed
     */
    async feedsIdGet(requestParameters: FeedsIdGetRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ServerFeedResponse> {
        const response = await this.feedsIdGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * update feed
     * update feed
     */
    async feedsIdPostRaw(requestParameters: FeedsIdPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ServerFeedUpdateResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling feedsIdPost().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/feeds/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ServerFeedUpdateResponseFromJSON(jsonValue));
    }

    /**
     * update feed
     * update feed
     */
    async feedsIdPost(requestParameters: FeedsIdPostRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ServerFeedUpdateResponse> {
        const response = await this.feedsIdPostRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * create feed
     * create feed
     */
    async feedsPostRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ServerFeedCreateResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/feeds`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ServerFeedCreateResponseFromJSON(jsonValue));
    }

    /**
     * create feed
     * create feed
     */
    async feedsPost(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ServerFeedCreateResponse> {
        const response = await this.feedsPostRaw(initOverrides);
        return await response.value();
    }

}