/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntWifiEdges,
    EntWifiEdgesFromJSON,
    EntWifiEdgesFromJSONTyped,
    EntWifiEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntWifi
 */
export interface EntWifi {
    /**
     * 
     * @type {EntWifiEdges}
     * @memberof EntWifi
     */
    edges?: EntWifiEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntWifi
     */
    id?: number;
    /**
     * Wifiname holds the value of the "wifiname" field.
     * @type {string}
     * @memberof EntWifi
     */
    wifiname?: string;
}

export function EntWifiFromJSON(json: any): EntWifi {
    return EntWifiFromJSONTyped(json, false);
}

export function EntWifiFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntWifi {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntWifiEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'wifiname': !exists(json, 'wifiname') ? undefined : json['wifiname'],
    };
}

export function EntWifiToJSON(value?: EntWifi | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntWifiEdgesToJSON(value.edges),
        'id': value.id,
        'wifiname': value.wifiname,
    };
}


