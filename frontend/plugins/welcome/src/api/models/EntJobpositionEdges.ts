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
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntJobpositionEdges
 */
export interface EntJobpositionEdges {
    /**
     * Employees holds the value of the employees edge.
     * @type {Array<EntEmployee>}
     * @memberof EntJobpositionEdges
     */
    employees?: Array<EntEmployee>;
}

export function EntJobpositionEdgesFromJSON(json: any): EntJobpositionEdges {
    return EntJobpositionEdgesFromJSONTyped(json, false);
}

export function EntJobpositionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntJobpositionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employees': !exists(json, 'employees') ? undefined : ((json['employees'] as Array<any>).map(EntEmployeeFromJSON)),
    };
}

export function EntJobpositionEdgesToJSON(value?: EntJobpositionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employees': value.employees === undefined ? undefined : ((value.employees as Array<any>).map(EntEmployeeToJSON)),
    };
}


