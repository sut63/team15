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
    EntRoomEdges,
    EntRoomEdgesFromJSON,
    EntRoomEdgesFromJSONTyped,
    EntRoomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoom
 */
export interface EntRoom {
    /**
     * 
     * @type {EntRoomEdges}
     * @memberof EntRoom
     */
    edges?: EntRoomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRoom
     */
    id?: number;
    /**
     * Roomprice holds the value of the "roomprice" field.
     * @type {number}
     * @memberof EntRoom
     */
    roomprice?: number;
    /**
     * Roomtypename holds the value of the "roomtypename" field.
     * @type {string}
     * @memberof EntRoom
     */
    roomtypename?: string;
}

export function EntRoomFromJSON(json: any): EntRoom {
    return EntRoomFromJSONTyped(json, false);
}

export function EntRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntRoomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'roomprice': !exists(json, 'roomprice') ? undefined : json['roomprice'],
        'roomtypename': !exists(json, 'roomtypename') ? undefined : json['roomtypename'],
    };
}

export function EntRoomToJSON(value?: EntRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntRoomEdgesToJSON(value.edges),
        'id': value.id,
        'roomprice': value.roomprice,
        'roomtypename': value.roomtypename,
    };
}


