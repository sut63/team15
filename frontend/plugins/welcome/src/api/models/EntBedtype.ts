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
    EntBedtypeEdges,
    EntBedtypeEdgesFromJSON,
    EntBedtypeEdgesFromJSONTyped,
    EntBedtypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBedtype
 */
export interface EntBedtype {
    /**
     * Bedtypename holds the value of the "bedtypename" field.
     * @type {string}
     * @memberof EntBedtype
     */
    bedtypename?: string;
    /**
     * 
     * @type {EntBedtypeEdges}
     * @memberof EntBedtype
     */
    edges?: EntBedtypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBedtype
     */
    id?: number;
}

export function EntBedtypeFromJSON(json: any): EntBedtype {
    return EntBedtypeFromJSONTyped(json, false);
}

export function EntBedtypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBedtype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bedtypename': !exists(json, 'bedtypename') ? undefined : json['bedtypename'],
        'edges': !exists(json, 'edges') ? undefined : EntBedtypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntBedtypeToJSON(value?: EntBedtype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bedtypename': value.bedtypename,
        'edges': EntBedtypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}

