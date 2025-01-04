if (!ObjC.available) {
    throw new Error('Objective-C runtime is not loaded');
}

function getClassHierarchy(objcObject: ObjC.Object): ObjC.Object[] {
    const hierarchy = [objcObject];
    while (objcObject.$superClass) {
        objcObject = objcObject.$superClass;
        hierarchy.push(objcObject);
    }

    return hierarchy;
}

for (const [className, objcClass] of Object.entries(ObjC.classes)) {
    if (!className.startsWith('HTModel')) {
        continue;
    }

    const hierarchy = getClassHierarchy(objcClass);

    console.log(className, hierarchy);
}
