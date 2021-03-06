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
 * @interface EntStaytypeEdges
 */
export interface EntStaytypeEdges {
    /**
     * Roomdetails holds the value of the roomdetails edge.
     * @type {Array<EntRoomdetail>}
     * @memberof EntStaytypeEdges
     */
    roomdetails?: Array<EntRoomdetail>;
}

export function EntStaytypeEdgesFromJSON(json: any): EntStaytypeEdges {
    return EntStaytypeEdgesFromJSONTyped(json, false);
}

export function EntStaytypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStaytypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomdetails': !exists(json, 'roomdetails') ? undefined : ((json['roomdetails'] as Array<any>).map(EntRoomdetailFromJSON)),
    };
}

export function EntStaytypeEdgesToJSON(value?: EntStaytypeEdges | null): any {
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


