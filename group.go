// Copyright 2015 pyros2097. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package spike

import (
	"github.com/pyros2097/spike/g2d"
	"github.com/pyros2097/spike/math/shape"
	"github.com/pyros2097/spike/math/vector"
)

// type Group struct {
//  Actor
// }

// func (self *Group) AddActor() {
//  // self.Draw(batch, parentAlpha)
// }

// var hello []*Actor
// var gp *Group = &Group{}

// func init() {
//  hello = []*Actor{
//    &gp.Actor,
//    &Actor{},
//  }
// }

// type IAA interface {
//  GetX() float32
// }

// type GG struct {
//  x float32
// }

// func (self *GG) GetX() float32 {
//  return self.x
// }

// type HH struct {
//  GG
// }

// var zzz []IAA

// // We achieve polymorhphism through interfaces
// // As now I can have an array of IActors which can contains Actors, Groups, Widgets, UI etc..

// func InitZZZ() {
//  zzz = []IAA{
//    &GG{},
//    &HH{},
//  }
// }

var (
	tmp = vector.NewVector2Empty()
)

// 2D scene graph node that may contain other actors.
// Actors have a z-order equal to the order they were inserted into the group. Actors inserted later will be drawn on top of
// actors added earlier. Touch events that hit more than one actor are distributed to topmost actors first.
type Group struct {
	Actor
	Children                        []*Actor
	worldTransform                  *vector.Affine2
	computedTransform, oldTransform *vector.Matrix4
	transform                       bool
	cullingArea                     shape.Rectangle
}

func NewGroup() *Group {
	return &Group{
		transform:         true,
		worldTransform:    vector.NewAffine2Empty(),
		oldTransform:      vector.NewMatrix4Empty(),
		computedTransform: vector.NewMatrix4Empty(),
	}
}

func (self *Group) act(delta float32) {
	self.Actor.act(delta)
	// Actor[] actors = children.begin();
	// for (int i = 0, n = children.size; i < n; i++)
	//    actors[i].act(delta);
	// children.end();
}

// Draws the group and its children. The default implementation calls {@link #applyTransform(Batch, Matrix4)} if needed, then
// {@link #drawChildren(Batch, float)}, then {@link #resetTransform(Batch)} if needed.
func (self *Group) draw(batch g2d.Batch, parentAlpha float32) {
	// if (transform) applyTransform(batch, computeTransform());
	// drawChildren(batch, parentAlpha);
	// if (transform) resetTransform(batch);
}

//  // Draws all children. {@link #applyTransform(Batch, Matrix4)} should be called before and {@link #resetTransform(Batch)} after
//   * this method if {@link #setTransform(boolean) transform} is true. If {@link #setTransform(boolean) transform} is false these
//   * methods don't need to be called, children positions are temporarily offset by the group position when drawn. This method
//   * avoids drawing children completely outside the {@link #setCullingArea(Rectangle) culling area}, if set.
//  protected func (self *Group) drawChildren (Batch batch, float parentAlpha) {
//    parentAlpha *= this.color.a;
//    SnapshotArray<Actor> children = this.children;
//    Actor[] actors = children.begin();
//    Rectangle cullingArea = this.cullingArea;
//    if (cullingArea != null) {
//      // Draw children only if inside culling area.
//      float cullLeft = cullingArea.x;
//      float cullRight = cullLeft + cullingArea.width;
//      float cullBottom = cullingArea.y;
//      float cullTop = cullBottom + cullingArea.height;
//      if (transform) {
//        for (int i = 0, n = children.size; i < n; i++) {
//          Actor child = actors[i];
//          if (!child.isVisible()) continue;
//          float cx = child.x, cy = child.y;
//          if (cx <= cullRight && cy <= cullTop && cx + child.width >= cullLeft && cy + child.height >= cullBottom)
//            child.draw(batch, parentAlpha);
//        }
//      } else {
//        // No transform for this group, offset each child.
//        float offsetX = x, offsetY = y;
//        x = 0;
//        y = 0;
//        for (int i = 0, n = children.size; i < n; i++) {
//          Actor child = actors[i];
//          if (!child.isVisible()) continue;
//          float cx = child.x, cy = child.y;
//          if (cx <= cullRight && cy <= cullTop && cx + child.width >= cullLeft && cy + child.height >= cullBottom) {
//            child.x = cx + offsetX;
//            child.y = cy + offsetY;
//            child.draw(batch, parentAlpha);
//            child.x = cx;
//            child.y = cy;
//          }
//        }
//        x = offsetX;
//        y = offsetY;
//      }
//    } else {
//      // No culling, draw all children.
//      if (transform) {
//        for (int i = 0, n = children.size; i < n; i++) {
//          Actor child = actors[i];
//          if (!child.isVisible()) continue;
//          child.draw(batch, parentAlpha);
//        }
//      } else {
//        // No transform for this group, offset each child.
//        float offsetX = x, offsetY = y;
//        x = 0;
//        y = 0;
//        for (int i = 0, n = children.size; i < n; i++) {
//          Actor child = actors[i];
//          if (!child.isVisible()) continue;
//          float cx = child.x, cy = child.y;
//          child.x = cx + offsetX;
//          child.y = cy + offsetY;
//          child.draw(batch, parentAlpha);
//          child.x = cx;
//          child.y = cy;
//        }
//        x = offsetX;
//        y = offsetY;
//      }
//    }
//    children.end();
//  }

//  // Draws this actor's debug lines if {@link #getDebug()} is true and, regardless of {@link #getDebug()}, calls
//   * {@link Actor#drawDebug(ShapeRenderer)} on each child.
// func (self *Group) drawDebug (ShapeRenderer shapes) {
//    drawDebugBounds(shapes);
//    if (transform) applyTransform(shapes, computeTransform());
//    drawDebugChildren(shapes);
//    if (transform) resetTransform(shapes);
//  }

//  // Draws all children. {@link #applyTransform(Batch, Matrix4)} should be called before and {@link #resetTransform(Batch)} after
//   * this method if {@link #setTransform(boolean) transform} is true. If {@link #setTransform(boolean) transform} is false these
//   * methods don't need to be called, children positions are temporarily offset by the group position when drawn. This method
//   * avoids drawing children completely outside the {@link #setCullingArea(Rectangle) culling area}, if set.
//  protected void drawDebugChildren (ShapeRenderer shapes) {
//    SnapshotArray<Actor> children = this.children;
//    Actor[] actors = children.begin();
//    // No culling, draw all children.
//    if (transform) {
//      for (int i = 0, n = children.size; i < n; i++) {
//        Actor child = actors[i];
//        if (!child.isVisible()) continue;
//        if (!child.getDebug() && !(child instanceof Group)) continue;
//        child.drawDebug(shapes);
//      }
//      shapes.flush();
//    } else {
//      // No transform for this group, offset each child.
//      float offsetX = x, offsetY = y;
//      x = 0;
//      y = 0;
//      for (int i = 0, n = children.size; i < n; i++) {
//        Actor child = actors[i];
//        if (!child.isVisible()) continue;
//        if (!child.getDebug() && !(child instanceof Group)) continue;
//        float cx = child.x, cy = child.y;
//        child.x = cx + offsetX;
//        child.y = cy + offsetY;
//        child.drawDebug(shapes);
//        child.x = cx;
//        child.y = cy;
//      }
//      x = offsetX;
//      y = offsetY;
//    }
//    children.end();
//  }

//  // Returns the transform for this group's coordinate system.
//  protected Matrix4 computeTransform () {
//    Affine2 worldTransform = this.worldTransform;

//    float originX = this.originX;
//    float originY = this.originY;
//    float rotation = this.rotation;
//    float scaleX = this.scaleX;
//    float scaleY = this.scaleY;

//    worldTransform.setToTrnRotScl(x + originX, y + originY, rotation, scaleX, scaleY);
//    if (originX != 0 || originY != 0) worldTransform.translate(-originX, -originY);

//    // Find the first parent that transforms.
//    Group parentGroup = parent;
//    while (parentGroup != null) {
//      if (parentGroup.transform) break;
//      parentGroup = parentGroup.parent;
//    }
//    if (parentGroup != null) worldTransform.preMul(parentGroup.worldTransform);

//    computedTransform.set(worldTransform);
//    return computedTransform;
//  }

//  // Set the batch's transformation matrix, often with the result of {@link #computeTransform()}. Note this causes the batch to
//   * be flushed. {@link #resetTransform(Batch)} will restore the transform to what it was before this call.
//  protected void applyTransform (Batch batch, Matrix4 transform) {
//    oldTransform.set(batch.getTransformMatrix());
//    batch.setTransformMatrix(transform);
//  }

//  // Restores the batch transform to what it was before {@link #applyTransform(Batch, Matrix4)}. Note this causes the batch to be
//   * flushed.
//  protected void resetTransform (Batch batch) {
//    batch.setTransformMatrix(oldTransform);
//  }

//  // Set the shape renderer transformation matrix, often with the result of {@link #computeTransform()}. Note this causes the
//   * shape renderer to be flushed. {@link #resetTransform(ShapeRenderer)} will restore the transform to what it was before this
//   * call.
//  protected void applyTransform (ShapeRenderer shapes, Matrix4 transform) {
//    oldTransform.set(shapes.getTransformMatrix());
//    shapes.setTransformMatrix(transform);
//  }

//  // Restores the shape renderer transform to what it was before {@link #applyTransform(Batch, Matrix4)}. Note this causes the
//   * shape renderer to be flushed.
//  protected void resetTransform (ShapeRenderer shapes) {
//    shapes.setTransformMatrix(oldTransform);
//  }

//  // Children completely outside of this rectangle will not be drawn. This is only valid for use with unrotated and unscaled
//   * actors!
// func (self *Group) setCullingArea (Rectangle cullingArea) {
//    this.cullingArea = cullingArea;
//  }

//  // @see #setCullingArea(Rectangle)
//  public Rectangle getCullingArea () {
//    return cullingArea;
//  }

//  public Actor hit (float x, float y, boolean touchable) {
//    if (touchable && getTouchable() == Touchable.disabled) return null;
//    Vector2 point = tmp;
//    Actor[] childrenArray = children.items;
//    for (int i = children.size - 1; i >= 0; i--) {
//      Actor child = childrenArray[i];
//      if (!child.isVisible()) continue;
//      child.parentToLocalCoordinates(point.set(x, y));
//      Actor hit = child.hit(point.x, point.y, touchable);
//      if (hit != null) return hit;
//    }
//    return super.hit(x, y, touchable);
//  }

//  // Called when actors are added to or removed from the group.
//  protected void childrenChanged () {
//  }

//  // Adds an actor as a child of this group. The actor is first removed from its parent group, if any.
// func (self *Group) addActor (Actor actor) {
//    if (actor.parent != null) actor.parent.removeActor(actor, false);
//    children.add(actor);
//    actor.setParent(this);
//    actor.setStage(getStage());
//    childrenChanged();
//  }

//  // Adds an actor as a child of this group, at a specific index. The actor is first removed from its parent group, if any.
//   * @param index May be greater than the number of children.
// func (self *Group) addActorAt (int index, Actor actor) {
//    if (actor.parent != null) actor.parent.removeActor(actor, false);
//    if (index >= children.size)
//      children.add(actor);
//    else
//      children.insert(index, actor);
//    actor.setParent(this);
//    actor.setStage(getStage());
//    childrenChanged();
//  }

//  // Adds an actor as a child of this group, immediately before another child actor. The actor is first removed from its parent
//   * group, if any.
// func (self *Group) addActorBefore (Actor actorBefore, Actor actor) {
//    if (actor.parent != null) actor.parent.removeActor(actor, false);
//    int index = children.indexOf(actorBefore, true);
//    children.insert(index, actor);
//    actor.setParent(this);
//    actor.setStage(getStage());
//    childrenChanged();
//  }

//  // Adds an actor as a child of this group, immediately after another child actor. The actor is first removed from its parent
//   * group, if any.
// func (self *Group) addActorAfter (Actor actorAfter, Actor actor) {
//    if (actor.parent != null) actor.parent.removeActor(actor, false);
//    int index = children.indexOf(actorAfter, true);
//    if (index == children.size)
//      children.add(actor);
//    else
//      children.insert(index + 1, actor);
//    actor.setParent(this);
//    actor.setStage(getStage());
//    childrenChanged();
//  }

//  // Calls {@link #removeActor(Actor, boolean)} with true.
//  public boolean removeActor (Actor actor) {
//    return removeActor(actor, true);
//  }

//  // Removes an actor from this group. If the actor will not be used again and has actions, they should be
//   * {@link Actor#clearActions() cleared} so the actions will be returned to their
//   * {@link Action#setPool(com.badlogic.gdx.utils.Pool) pool}, if any. This is not done automatically.
//   * @param unfocus If true, {@link Stage#unfocus(Actor)} is called.
//   * @return true if the actor was removed from this group.
//  public boolean removeActor (Actor actor, boolean unfocus) {
//    if (!children.removeValue(actor, true)) return false;
//    if (unfocus) {
//      Stage stage = getStage();
//      if (stage != null) stage.unfocus(actor);
//    }
//    actor.setParent(null);
//    actor.setStage(null);
//    childrenChanged();
//    return true;
//  }

//  // Removes all actors from this group.
// func (self *Group) clearChildren () {
//    Actor[] actors = children.begin();
//    for (int i = 0, n = children.size; i < n; i++) {
//      Actor child = actors[i];
//      child.setStage(null);
//      child.setParent(null);
//    }
//    children.end();
//    children.clear();
//    childrenChanged();
//  }

//  // Removes all children, actions, and listeners from this group.
// func (self *Group) clear () {
//    super.clear();
//    clearChildren();
//  }

//  // Returns the first actor found with the specified name. Note this recursively compares the name of every actor in the group.
//  public <T extends Actor> T findActor (String name) {
//    Array<Actor> children = this.children;
//    for (int i = 0, n = children.size; i < n; i++)
//      if (name.equals(children.get(i).getName())) return (T)children.get(i);
//    for (int i = 0, n = children.size; i < n; i++) {
//      Actor child = children.get(i);
//      if (child instanceof Group) {
//        Actor actor = ((Group)child).findActor(name);
//        if (actor != null) return (T)actor;
//      }
//    }
//    return null;
//  }

//  // Swaps two actors by index. Returns false if the swap did not occur because the indexes were out of bounds.
//  public boolean swapActor (int first, int second) {
//    int maxIndex = children.size;
//    if (first < 0 || first >= maxIndex) return false;
//    if (second < 0 || second >= maxIndex) return false;
//    children.swap(first, second);
//    return true;
//  }

//  // Swaps two actors. Returns false if the swap did not occur because the actors are not children of this group.
//  public boolean swapActor (Actor first, Actor second) {
//    int firstIndex = children.indexOf(first, true);
//    int secondIndex = children.indexOf(second, true);
//    if (firstIndex == -1 || secondIndex == -1) return false;
//    children.swap(firstIndex, secondIndex);
//    return true;
//  }

//  // Returns an ordered list of child actors in this group.
//  public SnapshotArray<Actor> getChildren () {
//    return children;
//  }

//  public boolean hasChildren () {
//    return children.size > 0;
//  }

//  // When true (the default), the Batch is transformed so children are drawn in their parent's coordinate system. This has a
//   * performance impact because {@link Batch#flush()} must be done before and after the transform. If the actors in a group are
//   * not rotated or scaled, then the transform for the group can be set to false. In this case, each child's position will be
//   * offset by the group's position for drawing, causing the children to appear in the correct location even though the Batch has
//   * not been transformed.
// func (self *Group) setTransform (boolean transform) {
//    this.transform = transform;
//  }

//  public boolean isTransform () {
//    return transform;
//  }

//  // Converts coordinates for this group to those of a descendant actor. The descendant does not need to be a direct child.
//   * @throws IllegalArgumentException if the specified actor is not a descendant of this group.
//  public Vector2 localToDescendantCoordinates (Actor descendant, Vector2 localCoords) {
//    Group parent = descendant.parent;
//    if (parent == null) throw new IllegalArgumentException("Child is not a descendant: " + descendant);
//    // First convert to the actor's parent coordinates.
//    if (parent != this) localToDescendantCoordinates(parent, localCoords);
//    // Then from each parent down to the descendant.
//    descendant.parentToLocalCoordinates(localCoords);
//    return localCoords;
//  }

//  // If true, {@link #drawDebug(ShapeRenderer)} will be called for this group and, optionally, all children recursively.
// func (self *Group) setDebug (boolean enabled, boolean recursively) {
//    setDebug(enabled);
//    if (recursively) {
//      for (Actor child : children) {
//        if (child instanceof Group) {
//          ((Group)child).setDebug(enabled, recursively);
//        } else {
//          child.setDebug(enabled);
//        }
//      }
//    }
//  }

//  // Calls {@link #setDebug(boolean, boolean)} with {@code true, true}.
//  public Group debugAll () {
//    setDebug(true, true);
//    return this;
//  }

//  // Returns a description of the actor hierarchy, recursively.
//  func (self *Group) toString() string {
//    StringBuilder buffer = new StringBuilder(128);
//    toString(buffer, 1);
//    buffer.setLength(buffer.length() - 1);
//    return buffer.toString();
//  }

//  void toString (StringBuilder buffer, int indent) {
//    buffer.append(super.toString());
//    buffer.append('\n');

//    Actor[] actors = children.begin();
//    for (int i = 0, n = children.size; i < n; i++) {
//      for (int ii = 0; ii < indent; ii++)
//        buffer.append("|  ");
//      Actor actor = actors[i];
//      if (actor instanceof Group)
//        ((Group)actor).toString(buffer, indent + 1);
//      else {
//        buffer.append(actor);
//        buffer.append('\n');
//      }
//    }
//    children.end();
//  }