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
    EntCleaningRoomEdges,
    EntCleaningRoomEdgesFromJSON,
    EntCleaningRoomEdgesFromJSONTyped,
    EntCleaningRoomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCleaningRoom
 */
export interface EntCleaningRoom {
    /**
     * Dateandstarttime holds the value of the "dateandstarttime" field.
     * @type {string}
     * @memberof EntCleaningRoom
     */
    dateandstarttime?: string;
    /**
     * 
     * @type {EntCleaningRoomEdges}
     * @memberof EntCleaningRoom
     */
    edges?: EntCleaningRoomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCleaningRoom
     */
    id?: number;
    /**
     * Note holds the value of the "note" field.
     * @type {string}
     * @memberof EntCleaningRoom
     */
    note?: string;
}

export function EntCleaningRoomFromJSON(json: any): EntCleaningRoom {
    return EntCleaningRoomFromJSONTyped(json, false);
}

export function EntCleaningRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCleaningRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dateandstarttime': !exists(json, 'dateandstarttime') ? undefined : json['dateandstarttime'],
        'edges': !exists(json, 'edges') ? undefined : EntCleaningRoomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'note': !exists(json, 'note') ? undefined : json['note'],
    };
}

export function EntCleaningRoomToJSON(value?: EntCleaningRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dateandstarttime': value.dateandstarttime,
        'edges': EntCleaningRoomEdgesToJSON(value.edges),
        'id': value.id,
        'note': value.note,
    };
}


