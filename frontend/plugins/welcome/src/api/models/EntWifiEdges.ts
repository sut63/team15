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
    EntLease,
    EntLeaseFromJSON,
    EntLeaseFromJSONTyped,
    EntLeaseToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntWifiEdges
 */
export interface EntWifiEdges {
    /**
     * Wifis holds the value of the wifis edge.
     * @type {Array<EntLease>}
     * @memberof EntWifiEdges
     */
    wifis?: Array<EntLease>;
}

export function EntWifiEdgesFromJSON(json: any): EntWifiEdges {
    return EntWifiEdgesFromJSONTyped(json, false);
}

export function EntWifiEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntWifiEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'wifis': !exists(json, 'wifis') ? undefined : ((json['wifis'] as Array<any>).map(EntLeaseFromJSON)),
    };
}

export function EntWifiEdgesToJSON(value?: EntWifiEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'wifis': value.wifis === undefined ? undefined : ((value.wifis as Array<any>).map(EntLeaseToJSON)),
    };
}


