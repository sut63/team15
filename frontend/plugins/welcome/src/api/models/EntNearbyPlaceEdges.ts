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
 * @interface EntNearbyplaceEdges
 */
export interface EntNearbyplaceEdges {
    /**
     * Roomdetail holds the value of the roomdetail edge.
     * @type {Array<EntRoomdetail>}
     * @memberof EntNearbyplaceEdges
     */
    roomdetail?: Array<EntRoomdetail>;
}

export function EntNearbyplaceEdgesFromJSON(json: any): EntNearbyplaceEdges {
    return EntNearbyplaceEdgesFromJSONTyped(json, false);
}

export function EntNearbyplaceEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntNearbyplaceEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomdetail': !exists(json, 'roomdetail') ? undefined : ((json['roomdetail'] as Array<any>).map(EntRoomdetailFromJSON)),
    };
}

export function EntNearbyplaceEdgesToJSON(value?: EntNearbyplaceEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomdetail': value.roomdetail === undefined ? undefined : ((value.roomdetail as Array<any>).map(EntRoomdetailToJSON)),
    };
}


