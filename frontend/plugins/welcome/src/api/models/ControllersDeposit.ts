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
/**
 * 
 * @export
 * @interface ControllersDeposit
 */
export interface ControllersDeposit {
    /**
     * 
     * @type {string}
     * @memberof ControllersDeposit
     */
    added?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersDeposit
     */
    depositorname?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersDeposit
     */
    depositortell?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDeposit
     */
    employee?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersDeposit
     */
    info?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDeposit
     */
    lease?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersDeposit
     */
    parcelcode?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersDeposit
     */
    recipienttell?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDeposit
     */
    statusd?: number;
}

export function ControllersDepositFromJSON(json: any): ControllersDeposit {
    return ControllersDepositFromJSONTyped(json, false);
}

export function ControllersDepositFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersDeposit {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'added': !exists(json, 'added') ? undefined : json['added'],
        'depositorname': !exists(json, 'depositorname') ? undefined : json['depositorname'],
        'depositortell': !exists(json, 'depositortell') ? undefined : json['depositortell'],
        'employee': !exists(json, 'employee') ? undefined : json['employee'],
        'info': !exists(json, 'info') ? undefined : json['info'],
        'lease': !exists(json, 'lease') ? undefined : json['lease'],
        'parcelcode': !exists(json, 'parcelcode') ? undefined : json['parcelcode'],
        'recipienttell': !exists(json, 'recipienttell') ? undefined : json['recipienttell'],
        'statusd': !exists(json, 'statusd') ? undefined : json['statusd'],
    };
}

export function ControllersDepositToJSON(value?: ControllersDeposit | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added': value.added,
        'depositorname': value.depositorname,
        'depositortell': value.depositortell,
        'employee': value.employee,
        'info': value.info,
        'lease': value.lease,
        'parcelcode': value.parcelcode,
        'recipienttell': value.recipienttell,
        'statusd': value.statusd,
    };
}


