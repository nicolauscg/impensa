import * as R from "ramda";

// source https://github.com/ramda/ramda/wiki/Cookbook#make-an-object-out-of-a-list-with-keys-derived-form-each-element
export const objFromListWith = R.curry((fn, list) =>
  R.chain(R.zipObj, R.map(fn))(list)
);

export const cleanNilFromObject = R.pickBy(
  R.pipe(
    R.isNil,
    R.not
  )
);

export const cleanEmptyFromObject = R.pickBy(
  R.pipe(
    R.isEmpty,
    R.not
  )
);

export const transformValuesToUpdateIdsPayload = values => {
  // eslint-disable-next-line no-unused-vars
  const { id, ...valuesWithoutId } = values;

  return {
    ids: Array.of(values.id),
    update: valuesWithoutId
  };
};

export const transformValuesToUpdateIdPayload = values => {
  // eslint-disable-next-line no-unused-vars
  const { id, ...valuesWithoutId } = values;

  return {
    id: values.id,
    update: valuesWithoutId
  };
};

export const pngImagetoBase64 = file =>
  new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
