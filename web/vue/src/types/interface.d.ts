declare module "Common" {
  export interface NameType {
    first: string;
    last: string;
  }

  export interface CardItemType {
    num: number;
    word: string;
  }
}

declare module "NmsCmp" {
  export interface SizeLimitType {
    size: number;
    unit: string;
  }

  export interface FormType {
    scoreThr: number;
    iouThr: number;
  }
}

declare module "Fetch" {
  export interface RespType {
    code: number;
    msg: string;
    data: any;
  }
}
