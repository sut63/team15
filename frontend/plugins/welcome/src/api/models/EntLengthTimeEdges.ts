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
    EntCleaningRoom,
    EntCleaningRoomFromJSON,
    EntCleaningRoomFromJSONTyped,
    EntCleaningRoomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntLengthTimeEdges
 */
export interface EntLengthTimeEdges {
    /**
     * Cleaningrooms holds the value of the cleaningrooms edge.
     * @type {Array<EntCleaningRoom>}
     * @memberof EntLengthTimeEdges
     */
    cleaningrooms?: Array<EntCleaningRoom>;
}

export function EntLengthTimeEdgesFromJSON(json: any): EntLengthTimeEdges {
    return EntLengthTimeEdgesFromJSONTyped(json, false);
}

export function EntLengthTimeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntLengthTimeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'cleaningrooms': !exists(json, 'cleaningrooms') ? undefined : ((json['cleaningrooms'] as Array<any>).map(EntCleaningRoomFromJSON)),
    };
}

export function EntLengthTimeEdgesToJSON(value?: EntLengthTimeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'cleaningrooms': value.cleaningrooms === undefined ? undefined : ((value.cleaningrooms as Array<any>).map(EntCleaningRoomToJSON)),
    };
}


