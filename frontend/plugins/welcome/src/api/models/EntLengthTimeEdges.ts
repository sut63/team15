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
    EntCleaningroom,
    EntCleaningroomFromJSON,
    EntCleaningroomFromJSONTyped,
    EntCleaningroomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntLengthtimeEdges
 */
export interface EntLengthtimeEdges {
    /**
     * Cleaningrooms holds the value of the cleaningrooms edge.
     * @type {Array<EntCleaningroom>}
     * @memberof EntLengthtimeEdges
     */
    cleaningrooms?: Array<EntCleaningroom>;
}

export function EntLengthtimeEdgesFromJSON(json: any): EntLengthtimeEdges {
    return EntLengthtimeEdgesFromJSONTyped(json, false);
}

export function EntLengthtimeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntLengthtimeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'cleaningrooms': !exists(json, 'cleaningrooms') ? undefined : ((json['cleaningrooms'] as Array<any>).map(EntCleaningroomFromJSON)),
    };
}

export function EntLengthtimeEdgesToJSON(value?: EntLengthtimeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'cleaningrooms': value.cleaningrooms === undefined ? undefined : ((value.cleaningrooms as Array<any>).map(EntCleaningroomToJSON)),
    };
}


