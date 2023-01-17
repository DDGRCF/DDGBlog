declare module "NmsCmp" {
  export interface sizeLimitType {
    size: number;
    unit: string;
  }

  export interface formType {
    scoreThr: number;
    iouThr: number;
  }
}

declare module "Fetch" {
  export interface respType {
    code: number;
    msg: string;
    data: any;
  }
}
