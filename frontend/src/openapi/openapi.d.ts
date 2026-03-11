import type {
  OpenAPIClient,
  Parameters,
  UnknownParamsObject,
  OperationResponse,
  AxiosRequestConfig,
} from 'openapi-client-axios';

declare namespace Components {
    namespace Parameters {
        /**
         * example:
         * -created_at
         */
        export type Sort = string;
    }
    export interface QueryParameters {
        sort?: /**
         * example:
         * -created_at
         */
        Parameters.Sort;
    }
    namespace Responses {
        export type LoginResponse = Schemas.User;
        export type NotFoundError = Schemas.Error;
        export interface PaymentListResponse {
            payments?: Schemas.Payment[];
        }
        export interface PaymentSummaryResponse {
            /**
             * example:
             * 40
             */
            total: number;
            /**
             * example:
             * 33
             */
            completed: number;
            /**
             * example:
             * 5
             */
            processing: number;
            /**
             * example:
             * 2
             */
            failed: number;
        }
        export type UnauthorizedError = Schemas.Error;
    }
    namespace Schemas {
        export interface Error {
            /**
             * example:
             * 401
             */
            code: number;
            /**
             * example:
             * Unauthenticated: missing or invalid token
             */
            message: string;
        }
        export interface Payment {
            /**
             * example:
             * 1
             */
            id: string;
            /**
             * example:
             * Merchant A
             */
            merchant: string;
            /**
             * example:
             * completed , processing , or failed
             */
            status: string;
            /**
             * example:
             * alice@example.com
             */
            amount: string;
            created_at: string; // date-time
        }
        export interface User {
            email: string;
            role: string;
            token: string;
        }
    }
}
declare namespace Paths {
    namespace DashboardV1AuthLogin {
        namespace Post {
            export interface RequestBody {
                email: string;
                password: string;
            }
            namespace Responses {
                export type $200 = Components.Responses.LoginResponse;
                export type $401 = Components.Responses.UnauthorizedError;
            }
        }
    }
    namespace DashboardV1Payments {
        namespace Get {
            namespace Parameters {
                export type Id = string;
                /**
                 * example:
                 * -created_at
                 */
                export type Sort = string;
                export type Status = string;
            }
            export interface QueryParameters {
                sort?: /**
                 * example:
                 * -created_at
                 */
                Parameters.Sort;
                status?: Parameters.Status;
                id?: Parameters.Id;
            }
            namespace Responses {
                export type $200 = Components.Responses.PaymentListResponse;
                export type $401 = Components.Responses.UnauthorizedError;
                export type $404 = Components.Responses.NotFoundError;
            }
        }
    }
    namespace DashboardV1PaymentsSummary {
        namespace Get {
            namespace Responses {
                export type $200 = Components.Responses.PaymentSummaryResponse;
                export type $401 = Components.Responses.UnauthorizedError;
            }
        }
    }
}


export interface OperationMethods {
}

export interface PathsDictionary {
  ['/dashboard/v1/auth/login']: {
  }
  ['/dashboard/v1/payments']: {
  }
  ['/dashboard/v1/payments/summary']: {
  }
}

export type Client = OpenAPIClient<OperationMethods, PathsDictionary>


export type Error = Components.Schemas.Error;
export type Payment = Components.Schemas.Payment;
export type User = Components.Schemas.User;
