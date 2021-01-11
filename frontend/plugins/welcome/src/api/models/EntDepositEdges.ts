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
    EntStatusd,
    EntStatusdFromJSON,
    EntStatusdFromJSONTyped,
    EntStatusdToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDepositEdges
 */
export interface EntDepositEdges {
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntDepositEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntStatusd}
     * @memberof EntDepositEdges
     */
    statusd?: EntStatusd;
}

export function EntDepositEdgesFromJSON(json: any): EntDepositEdges {
    return EntDepositEdgesFromJSONTyped(json, false);
}

export function EntDepositEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDepositEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'statusd': !exists(json, 'Statusd') ? undefined : EntStatusdFromJSON(json['Statusd']),
    };
}

export function EntDepositEdgesToJSON(value?: EntDepositEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': EntEmployeeToJSON(value.employee),
        'statusd': EntStatusdToJSON(value.statusd),
    };
}

