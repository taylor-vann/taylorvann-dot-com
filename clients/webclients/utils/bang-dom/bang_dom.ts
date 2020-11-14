// brian taylor vann

import { Structure } from "../bang/references/structure";
import { Bang } from "../bang/bang";
import { DocumentNode, AttributeKinds, hooks } from "./hooks/hooks";

const bang = new Bang(hooks);

const createStructure = <P, R>(
  structure: Structure<DocumentNode, AttributeKinds, P, R>
) => bang.createContextFactory(structure);

export { createStructure, bang };
