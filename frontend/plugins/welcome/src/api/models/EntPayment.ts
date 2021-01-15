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
    EntPaymentEdges,
    EntPaymentEdgesFromJSON,
    EntPaymentEdgesFromJSONTyped,
    EntPaymentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPayment
 */
export interface EntPayment {
    /**
     * 
     * @type {EntPaymentEdges}
     * @memberof EntPayment
     */
    edges?: EntPaymentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPayment
     */
    id?: number;
    /**
     * Paymentname holds the value of the "paymentname" field.
     * @type {string}
     * @memberof EntPayment
     */
    paymentname?: string;
}

export function EntPaymentFromJSON(json: any): EntPayment {
    return EntPaymentFromJSONTyped(json, false);
}

export function EntPaymentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPayment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntPaymentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'paymentname': !exists(json, 'paymentname') ? undefined : json['paymentname'],
    };
}

export function EntPaymentToJSON(value?: EntPayment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntPaymentEdgesToJSON(value.edges),
        'id': value.id,
        'paymentname': value.paymentname,
    };
}


