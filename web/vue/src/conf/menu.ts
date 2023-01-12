import { isString } from "@vueuse/shared";

interface menuItem {
  name: string;
  children?: menuItem[];
}

const menuObj = {
  elements: [
    {
      name: "旋转框处理工具",
      children: [
        {
          name: "DOTA提交格式转表格",
        },
        {
          name: "非极大值抑制合并",
        },
      ],
    },
  ] as menuItem[],
  getItem: getItemByIndex,
  getItemName: getItemNameByIndex,
};

function generateIndexs(): void {
  const menuElems = menuObj.elements;
  let menuElem: menuItem;
  const indexList = [];
  for (let i = 0; i < menuElems.length; i++) {
    menuElem = menuElems[i]; // 得到元素
    const index = String(i); // 得到元素的字符串下标
    indexList[i] = index; // 加入索引表
    if (menuElem.children) {
      indexList[i];
    }
  }
}

function getItemByIndex(
  _indexes: string | string[] | number[] | number
): menuItem | undefined {
  const indexList: number[] = [];

  if (Number.isInteger(_indexes)) {
    indexList[0] = _indexes as number;
  } else {
    if (isString(_indexes)) {
      _indexes = (_indexes as string).split("-");
    }
    for (let i = 0; i < (_indexes as (string | number)[]).length; ++i) {
      const num = Number((_indexes as (string | number)[])[i]);
      if (isNaN(num)) {
        return undefined;
      }
      indexList[i] = num;
    }
  }

  const elems: menuItem[] = menuObj.elements;
  let elem: menuItem = { name: "" };
  let i = 0;
  for (const index of indexList) {
    if (i == 0) {
      elem = elems[index];
      if (!elem) {
        return undefined;
      }
    } else {
      if (elem.children) {
        elem = elem.children![index];
        if (!elem) {
          return undefined;
        }
      } else {
        return undefined;
      }
    }
    i++;
  }
  return elem;
}

function getItemNameByIndex(
  _indexes: string | string[] | number[] | number
): string | undefined {
  return getItemByIndex(_indexes)?.name;
}

export default menuObj;
