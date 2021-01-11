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
    EntRoom,
    EntRoomFromJSON,
    EntRoomFromJSONTyped,
    EntRoomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFacilityEdges
 */
export interface EntFacilityEdges {
    /**
     * Room holds the value of the room edge.
     * @type {Array<EntRoom>}
     * @memberof EntFacilityEdges
     */
    room?: Array<EntRoom>;
}

export function EntFacilityEdgesFromJSON(json: any): EntFacilityEdges {
    return EntFacilityEdgesFromJSONTyped(json, false);
}

export function EntFacilityEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFacilityEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'room': !exists(json, 'room') ? undefined : ((json['room'] as Array<any>).map(EntRoomFromJSON)),
    };
}

export function EntFacilityEdgesToJSON(value?: EntFacilityEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'room': value.room === undefined ? undefined : ((value.room as Array<any>).map(EntRoomToJSON)),
    };
}

