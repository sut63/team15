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
 * @interface EntQuantityEdges
 */
export interface EntQuantityEdges {
    /**
     * Roomdetails holds the value of the roomdetails edge.
     * @type {Array<EntRoomdetail>}
     * @memberof EntQuantityEdges
     */
    roomdetails?: Array<EntRoomdetail>;
}

export function EntQuantityEdgesFromJSON(json: any): EntQuantityEdges {
    return EntQuantityEdgesFromJSONTyped(json, false);
}

export function EntQuantityEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntQuantityEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomdetails': !exists(json, 'Roomdetails') ? undefined : ((json['Roomdetails'] as Array<any>).map(EntRoomdetailFromJSON)),
    };
}

export function EntQuantityEdgesToJSON(value?: EntQuantityEdges | null): any {
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


