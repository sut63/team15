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
    EntRoomdetail,
    EntRoomdetailFromJSON,
    EntRoomdetailFromJSONTyped,
    EntRoomdetailToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBedtypeEdges
 */
export interface EntBedtypeEdges {
    /**
     * Roomdetails holds the value of the roomdetails edge.
     * @type {Array<EntRoomdetail>}
     * @memberof EntBedtypeEdges
     */
    roomdetails?: Array<EntRoomdetail>;
}

export function EntBedtypeEdgesFromJSON(json: any): EntBedtypeEdges {
    return EntBedtypeEdgesFromJSONTyped(json, false);
}

export function EntBedtypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBedtypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomdetails': !exists(json, 'roomdetails') ? undefined : ((json['roomdetails'] as Array<any>).map(EntRoomdetailFromJSON)),
    };
}

export function EntBedtypeEdgesToJSON(value?: EntBedtypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomdetails': value.roomdetails === undefined ? undefined : ((value.roomdetails as Array<any>).map(EntRoomdetailToJSON)),
    };
}

