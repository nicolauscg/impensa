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
