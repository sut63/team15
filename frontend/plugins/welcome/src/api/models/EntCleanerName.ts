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
    EntCleanerNameEdges,
    EntCleanerNameEdgesFromJSON,
    EntCleanerNameEdgesFromJSONTyped,
    EntCleanerNameEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCleanerName
 */
export interface EntCleanerName {
    /**
     * Cleanername holds the value of the "cleanername" field.
     * @type {string}
     * @memberof EntCleanerName
     */
    cleanername?: string;
    /**
     * 
     * @type {EntCleanerNameEdges}
     * @memberof EntCleanerName
     */
    edges?: EntCleanerNameEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCleanerName
     */
    id?: number;
}

export function EntCleanerNameFromJSON(json: any): EntCleanerName {
    return EntCleanerNameFromJSONTyped(json, false);
}

export function EntCleanerNameFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCleanerName {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'cleanername': !exists(json, 'cleanername') ? undefined : json['cleanername'],
        'edges': !exists(json, 'edges') ? undefined : EntCleanerNameEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntCleanerNameToJSON(value?: EntCleanerName | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'cleanername': value.cleanername,
        'edges': EntCleanerNameEdgesToJSON(value.edges),
        'id': value.id,
    };
}

