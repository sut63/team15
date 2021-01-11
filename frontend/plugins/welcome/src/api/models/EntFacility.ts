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
    EntFacilityEdges,
    EntFacilityEdgesFromJSON,
    EntFacilityEdgesFromJSONTyped,
    EntFacilityEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFacility
 */
export interface EntFacility {
    /**
     * 
     * @type {EntFacilityEdges}
     * @memberof EntFacility
     */
    edges?: EntFacilityEdges;
    /**
     * Facility holds the value of the "facility" field.
     * @type {string}
     * @memberof EntFacility
     */
    facility?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntFacility
     */
    id?: number;
}

export function EntFacilityFromJSON(json: any): EntFacility {
    return EntFacilityFromJSONTyped(json, false);
}

export function EntFacilityFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFacility {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntFacilityEdgesFromJSON(json['edges']),
        'facility': !exists(json, 'facility') ? undefined : json['facility'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntFacilityToJSON(value?: EntFacility | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntFacilityEdgesToJSON(value.edges),
        'facility': value.facility,
        'id': value.id,
    };
}

