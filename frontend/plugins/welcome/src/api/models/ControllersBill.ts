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
 * @interface ControllersBill
 */
export interface ControllersBill {
    /**
     * 
     * @type {string}
     * @memberof ControllersBill
     */
    added?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    payment?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    room?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    situation?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersBill
     */
    total?: number;
}

export function ControllersBillFromJSON(json: any): ControllersBill {
    return ControllersBillFromJSONTyped(json, false);
}

export function ControllersBillFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersBill {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'added': !exists(json, 'added') ? undefined : json['added'],
        'payment': !exists(json, 'payment') ? undefined : json['payment'],
        'room': !exists(json, 'room') ? undefined : json['room'],
        'situation': !exists(json, 'situation') ? undefined : json['situation'],
        'total': !exists(json, 'total') ? undefined : json['total'],
    };
}

export function ControllersBillToJSON(value?: ControllersBill | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'added': value.added,
        'payment': value.payment,
        'room': value.room,
        'situation': value.situation,
        'total': value.total,
    };
}

